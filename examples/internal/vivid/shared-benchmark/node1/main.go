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
		config.WithShared("127.0.0.1:8080")
	}))

	defer system1.Shutdown(true)

	ref := system1.ActorOfF(func() vivid.Actor {
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

	var cost time.Duration
	var n int
	go func() {
		msg := &Ping{}
		target := vivid.NewActorRef("127.0.0.1:8081", "/user/1")
		system1.FutureAsk(target, ref)
		startAt := time.Now()
		for i := 0; i < messageCount; i++ {
			system1.Tell(target, msg)
		}
		cost = time.Since(startAt)
		n = int(float32(messageCount*2) / (float32(cost) / float32(time.Second)))
		wait.Done()
	}()

	wait.Wait()
	log.Println("cost:", cost)
	log.Println("n:", n)
}
