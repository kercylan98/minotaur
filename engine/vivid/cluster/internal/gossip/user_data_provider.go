package gossip

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	gossipv1 "github.com/kercylan98/minotaur/engine/vivid/cluster/internal/gossip/v1"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"google.golang.org/protobuf/proto"
)

type UserDataProvider interface {
	Provide() (key string, value proto.Message)
}

type FunctionalUserDataProvider func() (key string, value proto.Message)

func (f FunctionalUserDataProvider) Provide() (key string, value proto.Message) {
	return f()
}

type StructUserDataProvider struct {
	Key   string
	Value proto.Message
}

func (s StructUserDataProvider) Provide() (key string, value proto.Message) {
	return s.Key, s.Value
}

// GetUserData 获取集群中的用户数据，当 nodeIds 不存在时，将获取所有节点
func GetUserData[M proto.Message](ctx vivid.ActorContext, gossipRef vivid.ActorRef, key string, nodeIds ...*NodeId) (map[gossipv1.NodeKey]M, error) {
	raw, err := vivid.FutureAsk[map[gossipv1.NodeKey]proto.Message](ctx, gossipRef, &GetUserDataRequestMessage{
		Key:     key,
		NodeIds: nodeIds,
	}).Result()
	if err != nil {
		return nil, err
	}

	if len(raw) == 0 {
		return make(map[gossipv1.NodeKey]M), nil
	}
	return collection.MappingFromMap(raw, func(value proto.Message) M {
		return value.(M)
	}), nil
}

// SetUserData 设置集群中的用户数据
func SetUserData(ctx vivid.ActorContext, gossipRef vivid.ActorRef, key string, value proto.Message) error {
	return vivid.FutureAsk[*gossipv1.GossipedAckMessage](ctx, gossipRef, &SetUserDataRequestMessage{
		Key:   key,
		Value: value,
	}).Wait()
}
