package vivid

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/internal/messages"
	"github.com/kercylan98/minotaur/toolkit/log"
	"sync"
)

// sharedSubscriptionStatusChangedMessage 用于通知集群节点状态变更的消息
type sharedSubscriptionStatusChangedMessage struct {
	Address prc.PhysicalAddress
	Closed  bool
}

func newSubscriptionActor(system *ActorSystem) *subscriptionActor {
	sa := &subscriptionActor{
		system: system,
	}
	sa.launched.Add(1)
	return sa
}

// subscriptionActor 用于管理发布与订阅的 Actor
type subscriptionActor struct {
	system     *ActorSystem                                // 所属 ActorSystem
	sas        map[prc.PhysicalAddress]ActorRef            // 远端 ActorSystem 的 subscriptionActor 的 ActorRef
	subscribes map[Topic]map[uint64]*messages.Subscription // 主题订阅者们，其中第二个 key 的寓意为订阅的 ID
	guid       uint64                                      // 订阅自增 ID 当前值
	ctx        context.Context                             // 订阅远端节点更新的上下文
	cancel     context.CancelFunc                          // 订阅远端节点更新的上下文取消函数，在调用后将不再订阅远端节点的更新
	launched   sync.WaitGroup                              // 意味着该 Actor 真实启动完成的等待信号
}

// bindSubscriptionContactProvider 由 ActorSystem 在创建该 Actor 时进行调用的绑定函数，它将监听集群中节点的变化
func (s *subscriptionActor) bindSubscriptionContactProvider(provider SubscriptionContactProvider) {
	go func() {
		s.launched.Wait()
		for {
			select {
			case <-s.ctx.Done():
				return
			case event := <-provider.ChangeNotify():
				_ = s.system.FutureAsk(s.system.subscription, &sharedSubscriptionStatusChangedMessage{Address: event.Address, Closed: event.Stop}).Wait()
			}
		}
	}()
}

func (s *subscriptionActor) OnReceive(ctx ActorContext) {
	switch m := ctx.Message().(type) {
	case *OnLaunch:
		s.sas = make(map[prc.PhysicalAddress]ActorRef)
		s.subscribes = make(map[Topic]map[uint64]*messages.Subscription)
		s.ctx, s.cancel = context.WithCancel(context.Background())
		s.launched.Done()
	case *OnTerminate:
		s.cancel()
	case *messages.SubscribeRequest:
		s.onSubscribeRequest(ctx, m)
	case *messages.UnsubscribeRequest:
		s.onUnsubscribeRequest(ctx, m)
	case *messages.PublishRequestBroadcast:
		s.onPublishRequestBroadcast(ctx, m)
	case *messages.LocalPublishRequest:
		s.onLocalPublishRequest(ctx, m)
	case *sharedSubscriptionStatusChangedMessage:
		s.onSharedSubscriptionStatusChangedMessage(ctx, m)
	}
}

func (s *subscriptionActor) onSubscribeRequest(ctx ActorContext, m *messages.SubscribeRequest) {
	subs, exist := s.subscribes[m.Topic]
	if !exist {
		subs = make(map[uint64]*messages.Subscription)
		s.subscribes[m.Topic] = subs
	}

	s.guid++
	sub := &messages.Subscription{
		Topic:      m.Topic,
		Id:         s.guid,
		Subscriber: m.Subscriber,
	}
	subs[sub.Id] = sub

	ctx.Reply(sub)
}

func (s *subscriptionActor) onUnsubscribeRequest(ctx ActorContext, m *messages.UnsubscribeRequest) {
	subs, exist := s.subscribes[m.Subscription.Topic]
	if !exist {
		return
	}

	delete(subs, m.Subscription.Id)
}

func (s *subscriptionActor) onPublishRequestBroadcast(ctx ActorContext, m *messages.PublishRequestBroadcast) {
	message, err := ctx.System().shared.GetCodec().Decode(m.MessageType, m.Data)
	if err != nil {
		ctx.System().Logger().Error("ActorSystem", log.String("type", "subscription"), log.Err(err))
		return
	}

	for _, subscription := range s.subscribes[m.Topic] {
		// 保持发送人
		ctx.(*actorContext).deliveryUserMessage(subscription.Subscriber, subscription.Subscriber, m.Publisher, nil, message)
		ctx.System().Logger().Debug("ActorSystem",
			log.String("system", ctx.System().PhysicalAddress()),
			log.String("type", "subscription"),
			log.String("type", "remote"),
			log.String("topic", m.Topic),
			log.String("publisher", m.Publisher.String()),
			log.String("subscriber", subscription.Subscriber.String()))
	}
}

func (s *subscriptionActor) onLocalPublishRequest(ctx ActorContext, m *messages.LocalPublishRequest) {
	if len(s.sas) > 0 {
		tn, data, err := ctx.System().shared.GetCodec().Encode(m.Message)
		if err != nil {
			panic(err)
		}

		broadcast := &messages.PublishRequestBroadcast{
			Data:        data,
			MessageType: tn,
			Publisher:   ctx.Sender(),
			Topic:       m.Topic,
		}

		for _, ref := range s.sas {
			ctx.Tell(ref, broadcast)
		}
	}

	for _, subscription := range s.subscribes[m.Topic] {
		// 保持发送人
		ctx.(*actorContext).deliveryUserMessage(subscription.Subscriber, subscription.Subscriber, ctx.Sender(), nil, m.Message)
		ctx.System().Logger().Debug("ActorSystem",
			log.String("system", ctx.System().PhysicalAddress()),
			log.String("type", "subscription"),
			log.String("type", "local"),
			log.String("topic", m.Topic),
			log.String("publisher", ctx.Sender().String()),
			log.String("subscriber", subscription.Subscriber.String()))
	}
}

func (s *subscriptionActor) onSharedSubscriptionStatusChangedMessage(ctx ActorContext, m *sharedSubscriptionStatusChangedMessage) {
	if m.Closed {
		delete(s.sas, m.Address)
	} else {
		s.sas[m.Address] = NewActorRef(m.Address, "/user/sub")
	}
	ctx.Reply(nil) // 告知完成
	ctx.System().Logger().Debug("ActorSystem", log.String("system", ctx.System().PhysicalAddress()), log.String("type", "subscription"), log.Any("changed", m))
}
