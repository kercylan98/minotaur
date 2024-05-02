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
	router.Route("system", t.onSystem)
	router.Register("account", "login").Bind(t.onAccountLogin)
}

func (t *TestService) testCall() {
	if err := t.router.Call("system")(123); err != nil {
		panic(err)
	}
	if err := t.router.Call("account", "login")(struct {
		Username, Password string
	}{Username: "username", Password: "pwd"}); err != nil {
		panic(err)
	}
}

func (t *TestService) onSystem(reader rpc.Reader) {
	var id int
	reader.ReadTo(&id)
	fmt.Println("call system", id)
}

func (t *TestService) onAccountLogin(reader rpc.Reader) {
	var params struct {
		Username string
		Password string
	}
	reader.ReadTo(&params)
	fmt.Println("call account login", params.Username, params.Password)
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
