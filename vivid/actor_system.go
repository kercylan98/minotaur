package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"golang.org/x/net/context"
	"reflect"
	"sync"
	"sync/atomic"
)

// NewActorSystem 创建一个 ActorSystem
func NewActorSystem(name string, opts ...*ActorSystemOptions) *ActorSystem {
	s := &ActorSystem{
		opts:         NewActorSystemOptions().Apply(opts...),
		name:         name,
		actors:       new(actorTrie).init(),
		closed:       make(chan struct{}),
		replyWaiters: make(map[uint64]chan any),
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())

	return s
}

// ActorSystem 是维护 Actor 的容器，负责 Actor 的创建、销毁、消息分发等
type ActorSystem struct {
	opts             *ActorSystemOptions
	name             string              // ActorSystem 的名称
	actors           *actorTrie          // Actor 容器
	closed           chan struct{}       // 关闭信号
	ctx              context.Context     // 上下文
	cancel           context.CancelFunc  // 取消函数（该函数应该在 closed 生效后才能调用）
	guid             atomic.Uint64       // 未命名 Actor 的唯一标识
	seq              atomic.Uint64       // 消息序列号
	replyWaiters     map[uint64]chan any // 等待回复的消息
	replyWaitersLock sync.Mutex          // 等待回复的消息锁
}

// Run 启动 ActorSystem
func (s *ActorSystem) Run() error {
	if s.opts.Host != "" {
		go s.handleRemoteMessage(s.ctx, s.opts.Server.C())
		return s.opts.Server.Run()
	}

	return nil
}

// ActorOf 创建一个 Actor
func ActorOf[T Actor](system *ActorSystem, opts ...*ActorOptions) (ActorRef, error) {
	typ := reflect.TypeOf((*T)(nil)).Elem()
	return system.ActorOf(typ, opts...)
}

// ActorOf 创建一个 Actor
//   - 推荐使用 ActorOf 函数来创建 Actor，这样可以保证 Actor 的类型安全
func (s *ActorSystem) ActorOf(typ reflect.Type, opts ...*ActorOptions) (ActorRef, error) {
	// 检查是否实现
	if !typ.Implements(actorType) {
		return nil, fmt.Errorf("%w: %s", ErrActorNotImplementActorRef, typ.String())
	}
	typ = typ.Elem()

	// 应用可选项
	opt := NewActorOptions().Apply(opts...)

	// 重复检查
	var actorId = NewActorId("tcp", s.opts.ClusterName, s.opts.Host, s.opts.Port, s.name, opt.Name)
	if opt.Name != charproc.None {
		if s.actors.has(actorId) {
			return nil, fmt.Errorf("%w: %s", ErrActorAlreadyExists, actorId.Name())
		}
	} else {
		opt.Name = fmt.Sprintf("%s-%d", typ.String(), s.guid.Add(1))
		actorId = NewActorId("tcp", s.opts.ClusterName, s.opts.Host, s.opts.Port, s.name, opt.Name)
	}

	// 创建 Actor
	var actor = reflect.New(typ).Interface().(Actor)
	var core = newActorCore(s, actorId, actor, opt)
	var dispatcher, exist = s.opts.Dispatchers[opt.DispatcherName]
	if !exist {
		dispatcher = s.opts.Dispatchers["default"]
	}
	if err := dispatcher.Attach(core); err != nil {
		return nil, err
	}
	if err := core.onPreStart(); err != nil {
		return nil, err
	}

	s.actors.insert(actorId, core)

	return core, nil
}

// GetActor 获取 ActorRef
func (s *ActorSystem) GetActor(actorId ActorId) (ActorRef, error) {
	if actorId.Cluster() != s.opts.ClusterName || actorId.Host() != s.opts.Host || actorId.Port() != s.opts.Port {
		// 生成远程 ActorRef
		return newRemoteActorRef(s, actorId), nil
	}

	if actorId.System() != s.name {
		return nil, fmt.Errorf("%w: %s", ErrActorNotFound, actorId.String())
	}

	ref := s.actors.find(actorId)
	if ref == nil {
		return nil, fmt.Errorf("%w: %s", ErrActorNotFound, actorId.String())
	}

	return ref, nil
}

func (s *ActorSystem) sendCtx(actorId ActorId, ctx MessageContext) error {
	core := s.actors.find(actorId)
	if core == nil {
		return fmt.Errorf("%w: %s", ErrActorNotFound, actorId.String())
	}
	receiverDispatcher, exist := s.opts.Dispatchers[core.GetOptions().DispatcherName]
	if !exist {
		receiverDispatcher = s.opts.Dispatchers["default"]
	}
	return receiverDispatcher.Send(core, ctx)
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

	// TODO: 客户端应该是一个连接池
	cli, err := s.opts.ClientFactory.NewClient(receiverId.Network(), receiverId.Host(), receiverId.Port())
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
			s.sendCtx(message.GetReceiverId(), message)
		}
	}
}
