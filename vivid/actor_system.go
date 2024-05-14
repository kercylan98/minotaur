package vivid

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/charproc"
	"golang.org/x/net/context"
	"reflect"
	"sync/atomic"
)

type actorSystemGetter func() *ActorSystem

// NewActorSystem 创建一个 ActorSystem
func NewActorSystem(name string, opts ...*ActorSystemOptions) *ActorSystem {
	s := &ActorSystem{
		opts:   NewActorSystemOptions().Apply(opts...),
		name:   name,
		actors: new(actorTrie).init(),
		closed: make(chan struct{}),
	}
	s.ctx, s.cancel = context.WithCancel(context.Background())

	return s
}

// ActorSystem 是维护 Actor 的容器，负责 Actor 的创建、销毁、消息分发等
type ActorSystem struct {
	opts   *ActorSystemOptions
	name   string             // ActorSystem 的名称
	actors *actorTrie         // Actor 容器
	closed chan struct{}      // 关闭信号
	ctx    context.Context    // 上下文
	cancel context.CancelFunc // 取消函数（该函数应该在 closed 生效后才能调用）
	guid   atomic.Uint64      // 未命名 Actor 的唯一标识
	seq    atomic.Uint64      // 消息序列号
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

	opt := NewActorOptions().Apply(opts...)

	// 重复检查
	var actorName = opt.Name
	if actorName != charproc.None {
		if s.actors.has(actorName) {
			return nil, fmt.Errorf("%w: %s", ErrActorAlreadyExists, actorName)
		}
	} else {
		actorName = fmt.Sprintf("%s-%d", typ.String(), s.guid.Add(1))
	}

	// 创建 Actor
	var imp = reflect.New(typ).Interface().(Actor)
	var a = new(localActor).init(
		opt,
		NewActorId("tcp", s.opts.ClusterName, s.opts.Host, s.opts.Port, s.name, actorName),
		imp,
		func() *ActorSystem {
			return s
		},
	)
	if err := imp.OnPreStart(a.ctx); err != nil {
		return nil, err
	}

	s.actors.insert(actorName, a)
	a.ctx.state = actorContextStateStarted

	// TODO: 临时使用普通异步循环处理消息，应该使用分发器来处理消息
	go func() {
		for message := range a.mailbox.Dequeue() {
			if err := imp.OnReceived(message); err != nil {
				fmt.Println(err)
			}
		}
	}()

	return a, nil
}

// GetActor 获取 ActorRef
func (s *ActorSystem) GetActor(actorId ActorId) (ActorRef, error) {
	if actorId.Cluster() != s.opts.ClusterName || actorId.Host() != s.opts.Host || actorId.Port() != s.opts.Port {
		// 生成远程 ActorRef
		return new(remoteActor).init(s, actorId), nil
	}

	if actorId.System() != s.name {
		return nil, fmt.Errorf("%w: %s", ErrActorNotFound, actorId.String())
	}

	ref := s.actors.find(actorId.Name())
	if ref == nil {
		return nil, fmt.Errorf("%w: %s", ErrActorNotFound, actorId.String())
	}

	return ref, nil
}

// send 用于向 Actor 发送消息
func (s *ActorSystem) send(receiver ActorRef, msg Message, opts ...MessageOption) ([]byte, error) {
	if receiver == nil {
		return nil, nil
	}
	opt := new(MessageOptions).apply(opts...)
	receiverId := receiver.GetId()

	// 检查是否为本地 Actor
	isLocal := receiverId.Host() == s.opts.Host && receiverId.Port() == s.opts.Port
	if isLocal {
		a := receiver.(*localActor)
		a.mailbox.Enqueue(msg) // TODO: 需兼容 Ask 模式
		return nil, nil
	}

	// 远程消息
	if s.opts.Host != "" {
		data, err := gob.Encode(newRemoteMessage(s, receiver, msg, opt))
		if err != nil {
			return nil, err
		}

		// TODO: 客户端应该是一个连接池
		cli, err := s.opts.ClientFactory.NewClient(receiverId.Network(), receiverId.Host(), receiverId.Port())
		if err != nil {
			return nil, err
		}

		var reply []byte
		if opt.Reply {
			reply, err = cli.Ask(data)
		} else {
			err = cli.Tell(data)
		}
		return reply, err
	}

	return nil, nil
}

// tell 用于向 Actor 发送消息
func (s *ActorSystem) tell(receiver ActorRef, msg Message, opts ...MessageOption) error {
	opt := new(MessageOptions).apply(append(opts, WithMessageReply(false))...)
	_, err := s.send(receiver, msg, WithMessageOptions(opt))
	return err
}

// ask 用于向 Actor 发送消息，并等待回复
func (s *ActorSystem) ask(receiver ActorRef, msg Message, opts ...MessageOption) ([]byte, error) {
	return s.send(receiver, msg, append(opts, WithMessageReply(true))...)
}

func (s *ActorSystem) handleRemoteMessage(ctx context.Context, c <-chan RemoteMessageEvent) {
	for {
		select {
		case <-ctx.Done():
			return
		case e := <-c:
			ref, err := s.GetActor(e.Message.Receiver)
			if err != nil {
				e.Callback(err)
				continue
			}

			err = s.tell(ref, e.Message.Message, WithMessageOptions(e.Message.Opts))
			e.Callback(err)
		}
	}
}
