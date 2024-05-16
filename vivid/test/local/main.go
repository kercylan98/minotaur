package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/chrono"
	"github.com/kercylan98/minotaur/toolkit/random"
	"github.com/kercylan98/minotaur/vivid"
	"time"
)

type LocalTestActor struct {
	vivid.BasicActor
	vm map[int]int
}

func (l *LocalTestActor) OnPreStart(ctx vivid.ActorContext) error {
	l.vm = make(map[int]int)
	return nil
}
func (l *LocalTestActor) OnReceived(ctx vivid.MessageContext) error {
	l.vm[ctx.GetMessage().(int)] += ctx.GetMessage().(int)
	fmt.Println(l.vm[ctx.GetMessage().(int)])
	return nil
}

func main() {
	system := vivid.NewActorSystem("LocalTest")
	if err := system.Run(); err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		ref, _ := vivid.ActorOf[*LocalTestActor](system)
		for j := 0; j < 1000; j++ {
			_ = ref.Tell(random.Int(0, 100))
		}
	}

	time.Sleep(time.Second * 5)
	if err := system.Shutdown(); err != nil {
		panic(err)
	}

	fmt.Println("Shutdown")
	time.Sleep(chrono.Week)
}
