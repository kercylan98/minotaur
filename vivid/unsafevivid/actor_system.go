package unsafevivid

import (
	"context"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"github.com/kercylan98/minotaur/toolkit/log"
	vivid "github.com/kercylan98/minotaur/vivid/vivids"
	"github.com/panjf2000/ants/v2"
	"path"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
)

var actorDispatcherContextKey = reflect.TypeOf((*vivid.Dispatcher)(nil)).Elem()

// NewActorSystem 创建一个 ActorSystem
func NewActorSystem(name string, opts ...*vivid.ActorSystemOptions) *ActorSystem {
	s := &ActorSystem{
		opts:         vivid.NewActorSystemOptions().Apply(opts...),
		name:         name,
		actors:       make(map[vivid.ActorId]*ActorCore),
		replyWaiters: make(map[uint64]chan any),
	}
	s.ActorSystemExternal = new(ActorSystemExternal).init(s)
	return s
}

// ActorSystem 是维护 Actor 的容器，负责 Actor 的创建、销毁、消息分发等
type ActorSystem struct {
	*ActorSystemExternal
	opts             *vivid.ActorSystemOptions    // ActorSystem 的配置项
	name             string                       // ActorSystem 的名称
	actors           map[vivid.ActorId]*ActorCore // 可用于精准快查的映射
	actorsRW         sync.RWMutex                 // 用于保护 actors 的读写锁
	user             *ActorCore                   // 用户使用的顶级 Actor
	ctx              context.Context              // 上下文
	cancel           context.CancelFunc           // 取消函数
	guid             atomic.Uint64                // 未命名 Actor 的唯一标识
	seq              atomic.Uint64                // 消息序列号
	replyWaiters     map[uint64]chan any          // 等待回复的消息
	replyWaitersLock sync.Mutex                   // 等待回复的消息锁
	gp               *ants.Pool                   // goroutine 池
}

// GetName 获取 ActorSystem 的名称
func (s *ActorSystem) GetName() string {
	return s.name
}

// Run 非阻塞的运行 ActorSystem
func (s *ActorSystem) Run() (err error) {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	pool, err := ants.NewPool(s.opts.AntsPoolSize, s.opts.AntsOptions...)
	if err != nil {
		return err
	}
	s.gp = pool

	defaultDispatcher, exist := s.opts.Dispatchers[""]
	if !exist {
		defaultDispatcher = NewDispatcher()
		s.opts.Dispatchers[""] = defaultDispatcher
	}
	for _, d := range s.opts.Dispatchers {
		d.OnInit(s)
	}

	s.user, err = s.generateActor(context.WithoutCancel(s.ctx), new(userGuardianActor), vivid.NewActorOptions().WithName("user"), false)
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
	s.releaseActor(s.user, true)
	delete(s.actors, s.user.GetId())

	for _, d := range s.opts.Dispatchers {
		d.Stop()
	}

	s.cancel()
	return nil
}

// ActorOf 创建一个 Actor
//   - 推荐使用 ActorOf 函数来创建 Actor，这样可以保证 Actor 的类型安全
func (s *ActorSystem) ActorOf(actor vivid.Actor, opts ...*vivid.ActorOptions) (vivid.ActorRef, error) {
	return s.user.ActorOf(actor, opts...)
}

// GetActor 获取 ActorRef
func (s *ActorSystem) GetActor() vivid.Query {
	return NewQuery(s, s.user)
}

func (s *ActorSystem) sendCtx(actorId vivid.ActorId, ctx vivid.MessageContext) error {
	ref, err := s.GetActor().MustActorId(actorId).One()
	if err != nil {
		return fmt.Errorf("%w: %s", err, actorId.String())
	}
	core := ref.(*ActorCore)

	return core.GetContext(actorDispatcherContextKey).(vivid.Dispatcher).Send(core, ctx)
}

// send 用于向 Actor 发送消息
func (s *ActorSystem) send(senderId, receiverId vivid.ActorId, msg vivid.Message, opts ...vivid.MessageOption) (vivid.MessageContext, *vivid.MessageOptions, error) {
	opt := new(vivid.MessageOptions).Apply(opts...)

	ctx := NewMessageContext(s, senderId, receiverId, msg, opt)

	// 检查是否为本地 Actor
	if isLocal := receiverId.Host() == s.opts.Host && receiverId.Port() == s.opts.Port; isLocal {
		return ctx, opt, s.sendCtx(receiverId, ctx)
	}

	// 远程消息如果是匿名发送，设置网络信息
	if ctx.GetSenderId() == "" {
		ctx.RemoteNetwork = s.opts.Network
		ctx.RemoteHost = s.opts.Host
		ctx.RemotePort = s.opts.Port
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

// Tell 用于向 Actor 发送消息
func (s *ActorSystem) Tell(receiverId vivid.ActorId, msg vivid.Message, opts ...vivid.MessageOption) error {
	opt := new(vivid.MessageOptions).Apply(append(opts, vivid.WithMessageReply(0))...)
	_, _, err := s.send(opt.SenderId, receiverId, msg, vivid.WithMessageOptions(opt))
	return err
}

// Ask 用于向 Actor 发送消息，并等待回复
func (s *ActorSystem) Ask(receiverId vivid.ActorId, msg vivid.Message, opts ...vivid.MessageOption) (any, error) {
	var opt = new(vivid.MessageOptions).Apply(opts...)
	var appendOpts = []vivid.MessageOption{vivid.WithMessageOptions(opt)}
	if opt.ReplyTimeout <= 0 {
		appendOpts = append(appendOpts, vivid.WithMessageReply(s.opts.AskDefaultTimeout))
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
		return nil, fmt.Errorf("%w: %s", vivid.ErrReplyTimeout, receiverId.String())
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
			reply, err := NewMessageContextWithBytes(s, data)
			if err != nil {
				panic(err)
			}

			// 处理回复消息
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
			if err = s.sendCtx(reply.GetReceiverId(), reply); err != nil {
				log.Error(fmt.Sprintf("handle remote message error: %s", err.Error()))
			}
		}
	}
}

func (s *ActorSystem) releaseActor(core *ActorCore, reenter bool) {
	if !reenter {
		s.actorsRW.RLock()
	}

	for key, child := range core.Children {
		if err := child.OnDestroy(child.Core); err != nil {
			log.Error(fmt.Sprintf("unregister actor destroy error: %s", err.Error()))
		}
		s.releaseActor(child, true)
		delete(core.Children, key)
		delete(s.actors, child.GetId())
	}

	delete(s.actors, core.Id)
	if core.Parent != nil {
		delete(core.Parent.Children, core.GetOptions().Name)
	}

	if !reenter {
		s.actorsRW.RUnlock()
	}

	if err := core.GetContext(actorDispatcherContextKey).(vivid.Dispatcher).Detach(core); err != nil {
		log.Error(fmt.Sprintf("unregister actor detach error: %s", err.Error()))
		return
	}

	log.Debug(fmt.Sprintf("actor %s released", core.GetId().String()))
}

func (s *ActorSystem) generateActor(ctx context.Context, actorImpl vivid.Actor, opt *vivid.ActorOptions, reenter bool) (*ActorCore, error) {
	// 应用可选项
	var actorPath string
	if opt.Name == "" {
		var builder strings.Builder
		builder.WriteString(opt.Parent.GetActorId().Name())
		builder.WriteByte('-')
		builder.WriteString(convert.Uint64ToString(s.guid.Add(1)))

		opt.Name = builder.String()
		if opt.Parent != nil {
			actorPath = path.Join(opt.Parent.GetActorId().Path(), opt.Name)
		} else {
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
	actorId := vivid.NewActorId(s.opts.Network, s.opts.ClusterName, s.opts.Host, s.opts.Port, s.name, actorPath)

	// 检查 Actor 是否已经存在
	if !reenter {
		s.actorsRW.Lock()
		defer s.actorsRW.Unlock()
	}
	actorCore, exist := s.actors[actorId]
	if exist {
		return nil, fmt.Errorf("%w: %s", vivid.ErrActorAlreadyExists, actorId.Name())
	}

	// 创建 Actor
	actorCore = NewActorCore(ctx, s, actorId, actorImpl, opt)

	// 绑定分发器
	dispatcher := s.opts.Dispatchers[opt.DispatcherName]
	if dispatcher == nil {
		return nil, fmt.Errorf("%w: %s", vivid.ErrDispatcherNotFound, opt.DispatcherName)
	}
	actorCore.SetContext(actorDispatcherContextKey, dispatcher)
	if err := dispatcher.Attach(actorCore); err != nil {
		s.releaseActor(actorCore, true)
		return nil, err
	}

	// 绑定父 Actor
	if opt.Parent != nil {
		actorCore.Parent = opt.Parent.(*ActorContext).Core
		opt.Parent.(*ActorContext).bindChildren(actorCore)
	}

	// 启动 Actor
	if err := actorCore.onPreStart(); err != nil {
		s.releaseActor(actorCore, true)
		return nil, err
	}

	s.actors[actorId] = actorCore
	log.Debug(fmt.Sprintf("actor %s created", actorId.String()))
	return actorCore, nil
}

// GetActorIds 获取 Actor ID
func (s *ActorSystem) GetActorIds() []vivid.ActorId {
	s.actorsRW.RLock()
	defer s.actorsRW.RUnlock()
	var ids = make([]vivid.ActorId, 0, len(s.actors))
	for _, actor := range s.actors {
		ids = append(ids, actor.GetId())
	}
	return ids
}
