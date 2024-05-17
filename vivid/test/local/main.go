package main

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/vivid"
	"time"
)

type LocalTestActorChild struct {
	vivid.BasicActor
	state int
}

func (l *LocalTestActorChild) OnPreStart(ctx vivid.ActorContext) error {

	return nil
}

func (l *LocalTestActorChild) OnReceived(ctx vivid.MessageContext) error {
	l.state++
	return nil
}

func (l *LocalTestActorChild) OnDestroy(ctx vivid.ActorContext) error {
	l.state = 0
	return nil
}

type LocalTestActor struct {
	vivid.BasicActor
}

func (l *LocalTestActor) OnPreStart(ctx vivid.ActorContext) error {
	return errors.New("test")
}

func (l *LocalTestActor) OnReceived(ctx vivid.MessageContext) error {
	time.Sleep(time.Second * 3)
	fmt.Println(ctx.GetMessage())
	return nil
}

func main() {
	system := vivid.NewActorSystem("LocalTest")
	if err := system.Run(); err != nil {
		panic(err)
	}

	ref, err := system.ActorOf(&LocalTestActor{}, nil)
	if err != nil {
		panic(err)
	}

	ref.Tell(123)

	if err := system.Shutdown(); err != nil {
		panic(err)
	}

	fmt.Println("Shutdown")
}
