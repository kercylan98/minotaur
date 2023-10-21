package lockstep_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/lockstep"
	"github.com/kercylan98/minotaur/utils/log"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
	"time"
)

type Cli struct {
	id string
}

func (slf *Cli) GetID() string {
	return slf.id
}

func (slf *Cli) Write(packet []byte, callback ...func(err error)) {
	log.Info("write", log.String("id", slf.id), log.String("frame", string(packet)))
}

func TestNewLockstep(t *testing.T) {
	ls := lockstep.NewLockstep[string, int]()
	ls.JoinClient(&Cli{id: "player_1"})
	ls.JoinClient(&Cli{id: "player_2"})
	count := 0
	ls.StartBroadcast()
	endChan := make(chan bool)

	go func() {
		for {
			ls.AddCommand(random.Int(1, 9999))
			count++
			if count >= 10 {
				break
			}
			time.Sleep(time.Millisecond * time.Duration(random.Int(10, 200)))
		}
		ls.StopBroadcast()
		endChan <- true
	}()

	<-endChan
	time.Sleep(time.Second)
	fmt.Println("end")
}
