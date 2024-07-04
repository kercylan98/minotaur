package vivid_test

import (
	"github.com/kercylan98/minotaur/core/vivid"
	"testing"
)

func TestOnReceiveFunc_OnReceive(t *testing.T) {
	called := false
	var f vivid.OnReceiveFunc = func(ctx vivid.ActorContext) {
		called = true
	}
	f.OnReceive(nil)

	if !called {
		t.Error("OnReceiveFunc.OnReceive() should call the function")
	}
}

func TestActorOnReceive(t *testing.T) {
	called := false
	var f vivid.Actor = vivid.FunctionalActor(func(ctx vivid.ActorContext) {
		called = true
	})
	f.OnReceive(nil)

	if !called {
		t.Error("Actor.OnReceive() should call the function")
	}
}
