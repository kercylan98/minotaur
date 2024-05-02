package rpccore_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/rpc"
	"github.com/kercylan98/minotaur/rpc/rpccore"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"testing"
	"time"
)

type TestService struct {
	router rpc.Router
}

func (t *TestService) OnRPCSetup(router rpc.Router) {
	t.router = router
	router.Route("account", t.onAccount)
}

func (t *TestService) testCall() {
	if err := t.router.Call("account")(123); err != nil {
		panic(err)
	}
}

func (t *TestService) onAccount(reader rpc.Reader) {
	var id int
	reader.ReadTo(&id)
	fmt.Println("call account", id)
}

func TestNats_OnServiceRegister(t *testing.T) {
	var nats, err = rpccore.NewNats("127.0.0.1:4222")
	if err != nil {
		panic(err)
	}

	ts := new(TestService)
	if err := rpc.NewApplication(nats).Register(
		ts,
	).Run(rpc.ServiceInfo{
		Name:     "test",
		UniqueId: "test-account",
	}); err != nil {
		panic(err)
	}

	go func() {
		for {
			time.Sleep(time.Second)
			ts.testCall()
		}
	}()

	time.Sleep(chrono.Day)
}
