package cmd

import (
	"github.com/kercylan98/minotaur/engine/prc"
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/spf13/cobra"
	"log"
	"sync"
	"time"
)

var (
	messageCount int
)

// benchmarkNetworkCmd represents the benchmarkNetwork command
var benchmarkNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "A command for benchmarking network transfer throughput",
	Long:  `It uses two Actors with Each behavior to detect the transfer rate`,
	Run: func(cmd *cobra.Command, args []string) {
		messageCount := messageCount
		wait := new(sync.WaitGroup)
		wait.Add(2)

		system1 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.WithShared("127.0.0.1:0")
		}))

		system2 := vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.WithShared("127.0.0.1:0")
		}))

		defer system1.Shutdown(true)
		defer system2.Shutdown(true)

		var ref1Sender, ref2Sender vivid.ActorRef

		ref1 := system1.ActorOfF(func() vivid.Actor {
			var count int

			return vivid.FunctionalActor(func(ctx vivid.ActorContext) {
				switch ctx.Message().(type) {
				case vivid.ActorRef:
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
				case vivid.ActorRef:
					ctx.Tell(ref2Sender, ctx.Ref())
				}
			})
		})
		ref1Sender = ref2.Clone()

		var cost time.Duration
		var n int
		go func() {
			msg := &prc.ProcessId{}
			startAt := time.Now()
			for i := 0; i < messageCount; i++ {
				system1.Tell(ref1Sender, msg)
			}
			cost = time.Since(startAt)
			n = int(float32(messageCount*2) / (float32(cost) / float32(time.Second)))
			wait.Done()
		}()

		wait.Wait()
		log.Println("sec:", n)
		log.Println("cost:", cost)
	},
}

func init() {
	benchmarkCmd.AddCommand(benchmarkNetworkCmd)

	benchmarkNetworkCmd.Flags().IntVarP(&messageCount, "message count", "c", 1000000, "message count")
}
