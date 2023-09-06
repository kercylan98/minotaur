package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/timer"
	"github.com/kercylan98/minotaur/utils/times"
	"sync"
)

func main() {
	var ticker = timer.GetTicker(10)
	var wait sync.WaitGroup
	wait.Add(3)
	ticker.Loop("LOOP", timer.Instantly, times.Second, timer.Forever, func() {
		fmt.Println("LOOP")
		wait.Done()
	})
	wait.Wait()
}
