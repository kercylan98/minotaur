package moduleimpl

import (
	"fmt"
	"github.com/kercylan98/minotaur/engine/future"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/engine/vivid/cluster"
	"github.com/kercylan98/minotaur/skeleton/internal/module"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/toolkit/collection"
	"google.golang.org/protobuf/proto"
	"sync"
	"time"
)

const (
	rpcClusterUserDataKey = "rpc_cluster_user_data"
)

var _ module.RPCModule = (*RPCModule)(nil)

func NewRPCModule() *RPCModule {
	return &RPCModule{}
}

type RPCModule struct {
	modules struct {
		actorSystem module.ActorSystemModule
	}

	userDataRW sync.RWMutex
	userData   *RPCModuleClusterUserData
}

func (r *RPCModule) OnInitialize(ctx *application.Context) (err error) {
	r.userData = &RPCModuleClusterUserData{
		Handlers: make(map[string]*RPCModuleClusterUserData_Handlers),
	}
	return
}

func (r *RPCModule) OnDependencyInitialize(ctx *application.Context) (err error) {
	r.modules.actorSystem = application.LoadModule[module.ActorSystemModule](ctx)
	return
}

func (r *RPCModule) OnDependencySetup() (err error) {
	if r.modules.actorSystem.Cluster() == nil {
		return fmt.Errorf("cluster mode is not enabled, please check the configuration [%s]", application.ConfigurationVividClusterSeedNodesKey)
	}
	return
}

func (r *RPCModule) OnSetup() (err error) {
	return
}

func (r *RPCModule) RegisterMessage(ctx vivid.ActorContext, message proto.Message) {
	messageTypeName := string(proto.MessageName(message))

	r.userDataRW.Lock()
	r.userData.Handlers[messageTypeName].Ref = append(r.userData.Handlers[messageTypeName].Ref, ctx.Ref())
	collection.DeduplicateSliceInPlaceWithCompare(&r.userData.Handlers[messageTypeName].Ref, func(a, b vivid.ActorRef) bool {
		return a.Equal(b)
	})
	r.userDataRW.Unlock()
}

func (r *RPCModule) AffirmMessage() error {
	r.userDataRW.RLock()
	clone := proto.Clone(r.userData).(*RPCModuleClusterUserData)
	r.userDataRW.RUnlock()

	if err := cluster.SetUserData(r.modules.actorSystem.Cluster(), rpcClusterUserDataKey, clone); err != nil {
		return err
	}
	return nil
}

func (r *RPCModule) Request(message proto.Message, timeout ...time.Duration) future.Future[vivid.Message] {

	nodes, err := cluster.GetUserData[*RPCModuleClusterUserData](r.modules.actorSystem.Cluster(), rpcClusterUserDataKey)
	if err != nil {
		return future.Fail[vivid.Message](err)
	}

	messageTypeName := string(proto.MessageName(message))

	for _, userData := range nodes {
		refs, exist := userData.Handlers[messageTypeName]
		if !exist {
			continue
		}

		selected := collection.ChooseRandomSliceElement(refs.Ref)
		return r.modules.actorSystem.Cluster().FutureAsk(selected, message, timeout...)
	}

	return future.Fail[vivid.Message](fmt.Errorf("no handler for message [%s]", messageTypeName))
}

func (r *RPCModule) Tell(message proto.Message) {
	nodes, err := cluster.GetUserData[*RPCModuleClusterUserData](r.modules.actorSystem.Cluster(), rpcClusterUserDataKey)
	if err != nil {
		return
	}

	messageTypeName := string(proto.MessageName(message))

	for _, userData := range nodes {
		refs, exist := userData.Handlers[messageTypeName]
		if !exist {
			continue
		}

		for _, ref := range refs.Ref {
			r.modules.actorSystem.Cluster().Tell(ref, message)
		}
	}
}
