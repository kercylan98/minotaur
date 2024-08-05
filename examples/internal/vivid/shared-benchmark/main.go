package main

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"log"
	"sync"
	"time"
)

//goland:noinspection t
func main() {
	messageCount := 1000000
	wait := new(sync.WaitGroup)
	wait.Add(2)

	system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared(":8080")
	}))

	system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
		config.WithShared(":8081")
	}))

	defer system1.Shutdown(true)
	defer system2.Shutdown(true)

	var ref1Sender, ref2Sender vivid.ActorRef

	ref1 := system1.ActorOfF(func() vivid.Actor {
		var count int

		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *Pong:
				count++
				if count%50000 == 0 {
					log.Println(count)
				}
				if count == messageCount {
					wait.Done()
				}
			}
		})
	})
	ref2Sender = ref1.Clone()

	ref2 := system2.ActorOfF(func() vivid.Actor {
		return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
			switch ctx.Message().(type) {
			case *Ping:
				ctx.Tell(ref2Sender, &Pong{})
			}
		})
	})
	ref1Sender = ref2.Clone()

	var cost time.Duration
	var n int
	go func() {
		msg := &Ping{}
		startAt := time.Now()
		for i := 0; i < messageCount; i++ {
			system1.Tell(ref1Sender, msg)
		}
		cost = time.Since(startAt)
		n = int(float32(messageCount*2) / (float32(cost) / float32(time.Second)))
		wait.Done()
	}()

	wait.Wait()
	log.Println("cost:", cost)
	log.Println("n:", n)
}
