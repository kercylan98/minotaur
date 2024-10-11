package vivid

import (
	"context"
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid/internal/messages"
	"github.com/kercylan98/minotaur/toolkit/log"
	"sync"
)

type sharedSubscriptionStatusChangedMessage struct {
	address prc.PhysicalAddress
	closed  bool
}

func newSubscriptionActor(system *ActorSystem) *subscriptionActor {
	sa := &subscriptionActor{
		system: system,
	}
	sa.launched.Add(1)
	return sa
}

type subscriptionActor struct {
	system     *ActorSystem
	sas        map[prc.PhysicalAddress]ActorRef
	subscribes map[Topic]map[uint64]*messages.Subscription
	guid       uint64
	ctx        context.Context
	cancel     context.CancelFunc
	launched   sync.WaitGroup
}

func (s *subscriptionActor) BindSubscriptionContactProvider(provider SubscriptionContactProvider) {
	go func() {
		s.launched.Wait()
		for {
			select {
			case <-s.ctx.Done():
				return
			case event := <-provider.ChangeNotify():
				s.system.Tell(s.system.subscription, &sharedSubscriptionStatusChangedMessage{address: event.Address, closed: event.Stop})
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
		ctx.(*actorContext).deliveryUserMessage(subscription.Subscriber, subscription.Subscriber, ctx.Sender(), nil, message)
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
	}
}

func (s *subscriptionActor) onSharedSubscriptionStatusChangedMessage(ctx ActorContext, m *sharedSubscriptionStatusChangedMessage) {
	if m.closed {
		delete(s.sas, m.address)
	} else {
		s.sas[m.address] = NewActorRef(m.address, "/user/sub")
	}
}
