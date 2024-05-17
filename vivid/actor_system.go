package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/log"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/net/context"
	"path"
	"sync"
	"sync/atomic"
)

// NewActorSystem 创建一个 ActorSystem
func NewActorSystem(name string, opts ...*ActorSystemOptions) *ActorSystem {
	s := &ActorSystem{
		opts:         NewActorSystemOptions().Apply(opts...),
		name:         name,
		actors:       make(map[ActorId]*actorCore),
		replyWaiters: make(map[uint64]chan any),
	}
	s.actorSystemExternal = new(actorSystemExternal).init(s)
	return s
}

// ActorSystem 是维护 Actor 的容器，负责 Actor 的创建、销毁、消息分发等
type ActorSystem struct {
	*actorSystemExternal
	opts             *ActorSystemOptions
	name             string                 // ActorSystem 的名称
	actors           map[ActorId]*actorCore // 可用于精准快查的映射
	actorsRW         sync.RWMutex           // 用于保护 actors 的读写锁
	user             *actorCore             // 用户使用的顶级 Actor
	ctx              context.Context        // 上下文
	cancel           context.CancelFunc     // 取消函数
	guid             atomic.Uint64          // 未命名 Actor 的唯一标识
	seq              atomic.Uint64          // 消息序列号
	replyWaiters     map[uint64]chan any    // 等待回复的消息
	replyWaitersLock sync.Mutex             // 等待回复的消息锁
	gp               *ants.Pool             // goroutine 池
}

// Run 非阻塞的运行 ActorSystem
func (s *ActorSystem) Run() (err error) {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	pool, err := ants.NewPool(s.opts.AntsPoolSize, s.opts.AntsOptions...)
	if err != nil {
		return err
	}
	s.gp = pool

	for _, d := range s.opts.Dispatchers {
		d.OnInit(s)
	}

	s.user, err = s.generateActor(new(userGuardianActor), NewActorOptions().WithName("user"))
	if err != nil {
		return err
	}

	if s.opts.Host != "" {
		for i := 0; i < int(s.opts.RemoteProcessorNum); i++ {
			go s.handleRemoteMessage(s.ctx, s.opts.Server.C())
		}
	}

	return nil
}

// Shutdown 关闭 ActorSystem
func (s *ActorSystem) Shutdown() error {
	if err := s.user.OnDestroy(s.user); err != nil {
		return err
	}
	s.actorsRW.Lock()
	defer s.actorsRW.Unlock()
	s.unregisterActor(s.user, true)
	delete(s.actors, s.user.GetId())

	for _, d := range s.opts.Dispatchers {
		d.Stop()
	}

	s.cancel()
	return nil
}

// ActorOf 创建一个 Actor
//   - 推荐使用 ActorOf 函数来创建 Actor，这样可以保证 Actor 的类型安全
func (s *ActorSystem) ActorOf(actor Actor, opts ...*ActorOptions) (ActorRef, error) {
	return s.user.ActorOf(actor, opts...)
}

// GetActor 获取 ActorRef
func (s *ActorSystem) GetActor() Query {
	return newQuery(s, s.user)
}

func (s *ActorSystem) getActorDispatcher(core ActorCore) Dispatcher {
	receiverDispatcher, exist := s.opts.Dispatchers[core.GetOptions().DispatcherName]
	if !exist {
		receiverDispatcher = s.opts.Dispatchers["default"]
	}
	return receiverDispatcher
}

func (s *ActorSystem) sendCtx(actorId ActorId, ctx MessageContext) error {
	core, err := s.GetActor().MustActorId(actorId).internalOne()
	if err != nil {
		return fmt.Errorf("%w: %s", err, actorId.String())
	}
	return s.getActorDispatcher(core).Send(core, ctx)
}

// send 用于向 Actor 发送消息
func (s *ActorSystem) send(senderId, receiverId ActorId, msg Message, opts ...MessageOption) (MessageContext, *MessageOptions, error) {
	opt := new(MessageOptions).apply(opts...)

	ctx := newMessageContext(s, senderId, receiverId, msg, opt)

	// 检查是否为本地 Actor
	if isLocal := receiverId.Host() == s.opts.Host && receiverId.Port() == s.opts.Port; isLocal {
		return ctx, opt, s.sendCtx(receiverId, ctx)
	}

	// 远程消息如果是匿名发送，设置网络信息
	if ctx.GetSenderId() == "" {
		c := ctx.(*messageContext)
		c.RemoteNetwork = s.opts.Network
		c.RemoteHost = s.opts.Host
		c.RemotePort = s.opts.Port
	}

	data, err := gob.Encode(ctx)
	if err != nil {
		return nil, nil, err
	}

	cli, err := s.opts.ClientFactory.NewClient(s, receiverId.Network(), receiverId.Host(), receiverId.Port())
	if err != nil {
		return nil, nil, err
	}

	return ctx, opt, cli.Exec(data)
}

// tell 用于向 Actor 发送消息
func (s *ActorSystem) tell(receiverId ActorId, msg Message, opts ...MessageOption) error {
	opt := new(MessageOptions).apply(append(opts, WithMessageReply(0))...)
	_, _, err := s.send(opt.SenderId, receiverId, msg, WithMessageOptions(opt))
	return err
}

// ask 用于向 Actor 发送消息，并等待回复
func (s *ActorSystem) ask(receiverId ActorId, msg Message, opts ...MessageOption) (any, error) {
	var opt = new(MessageOptions).apply(opts...)
	var appendOpts = []MessageOption{WithMessageOptions(opt)}
	if opt.ReplyTimeout <= 0 {
		appendOpts = append(appendOpts, WithMessageReply(s.opts.AskDefaultTimeout))
	}

	ctx, opt, err := s.send(opt.SenderId, receiverId, msg, appendOpts...)
	if err != nil {
		return nil, err
	}

	waiter := make(chan any)
	seq := ctx.GetSeq()
	s.replyWaitersLock.Lock()
	s.replyWaiters[seq] = waiter
	s.replyWaitersLock.Unlock()

	timeout, cancel := context.WithTimeout(s.ctx, opt.ReplyTimeout)
	defer func(s *ActorSystem, seq uint64, cancel context.CancelFunc, waiter chan any) {
		cancel()
		close(waiter)
		s.replyWaitersLock.Lock()
		delete(s.replyWaiters, seq)
		s.replyWaitersLock.Unlock()
	}(s, seq, cancel, waiter)

	select {
	case <-timeout.Done():
		return nil, fmt.Errorf("%w: %s", ErrReplyTimeout, receiverId.String())
	case reply := <-waiter:
		err, ok := reply.(error)
		if ok {
			return nil, err
		}
		return reply, nil
	}
}

func (s *ActorSystem) handleRemoteMessage(ctx context.Context, c <-chan []byte) {
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-c:
			message, err := parseMessageContext(s, data)
			if err != nil {
				panic(err)
			}

			// 处理回复消息
			reply := message.(*messageContext)
			if reply.ReplyMessage != nil {
				s.replyWaitersLock.Lock()
				waiter, exist := s.replyWaiters[reply.Seq]
				s.replyWaitersLock.Unlock()
				if exist {
					waiter <- reply.ReplyMessage
				}
				continue
			}

			// 处理请求消息
			if err = s.sendCtx(message.GetReceiverId(), message); err != nil {
				log.Error(fmt.Sprintf("handle remote message error: %s", err.Error()))
			}
		}
	}
}

func (s *ActorSystem) unregisterActor(core *actorCore, reenter bool) {
	if !reenter {
		s.actorsRW.RLock()
	}

	for key, child := range core.children {
		if err := child.OnDestroy(child.core); err != nil {
			log.Error(fmt.Sprintf("unregister actor destroy error: %s", err.Error()))
		}
		s.unregisterActor(child, true)
		delete(core.children, key)
		delete(s.actors, child.GetId())
	}

	delete(s.actors, core.id)
	if core.parent != nil {
		delete(core.parent.children, core.GetOptions().Name)
	}

	if !reenter {
		s.actorsRW.RUnlock()
	}

	if err := s.getActorDispatcher(core).Detach(core); err != nil {
		log.Error(fmt.Sprintf("unregister actor detach error: %s", err.Error()))
		return
	}
}

func (s *ActorSystem) generateActor(actorImpl Actor, opts ...*ActorOptions) (*actorCore, error) {
	// 应用可选项
	opt := NewActorOptions().Apply(opts...)

	var actorPath string
	if opt.Name == "" {
		if opt.Parent != nil {
			opt.Name = fmt.Sprintf("%s-%d", opt.Parent.GetActorId().Name(), s.guid.Add(1))
			actorPath = path.Join(opt.Parent.GetActorId().Path(), opt.Name)
		} else {
			opt.Name = fmt.Sprintf("%s-%d", s.name, s.guid.Add(1))
			actorPath = path.Clean(opt.Name)
		}
	} else {
		if opt.Parent != nil {
			actorPath = path.Join(opt.Parent.GetActorId().Path(), opt.Name)
		} else {
			actorPath = path.Clean(opt.Name)
		}
	}

	// 创建 Actor ID
	actorId := NewActorId(s.opts.Network, s.opts.ClusterName, s.opts.Host, s.opts.Port, s.name, actorPath)

	// 检查 Actor 是否已经存在
	s.actorsRW.Lock()
	defer s.actorsRW.Unlock()
	actor, exist := s.actors[actorId]
	if exist {
		return nil, fmt.Errorf("%w: %s", ErrActorAlreadyExists, actorId.Name())
	}

	// 创建 Actor
	actor = newActorCore(s, actorId, actorImpl, opt)

	// 绑定分发器
	if err := s.getActorDispatcher(actor).Attach(actor); err != nil {
		s.unregisterActor(actor, true)
		return nil, err
	}

	// 绑定父 Actor
	if opt.Parent != nil {
		actor.parent = opt.Parent.(*actorContext).core
		opt.Parent.(*actorContext).bindChildren(actor)
	}

	// 启动 Actor
	if err := actor.onPreStart(); err != nil {
		s.unregisterActor(actor, true)
		return nil, err
	}

	s.actors[actorId] = actor
	return actor, nil
}

// GetActorIds 获取 Actor ID
func (s *ActorSystem) GetActorIds() []ActorId {
	s.actorsRW.RLock()
	defer s.actorsRW.RUnlock()
	var ids = make([]ActorId, 0, len(s.actors))
	for _, actor := range s.actors {
		ids = append(ids, actor.GetId())
	}
	return ids
}

// GetName 获取 ActorSystem 的名称
func (s *ActorSystem) GetName() string {
	return s.name
}
