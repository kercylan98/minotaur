package moving_test

import (
	"fmt"
	moving2 "github.com/kercylan98/minotaur/utils/moving"
	"sync"
	"time"
)

func ExampleNewTwoDimensional() {
	m := moving2.NewTwoDimensional[int64, float64]()
	defer func() {
		m.Release()
	}()
	fmt.Println(m != nil)

	// Output:
	// true
}

func ExampleTwoDimensional_MoveTo() {
	m := moving2.NewTwoDimensional(moving2.WithTwoDimensionalTimeUnit[int64, float64](time.Second))
	defer func() {
		m.Release()
	}()

	var wait sync.WaitGroup
	m.RegPosition2DDestinationEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println("done")
		wait.Done()
	})

	wait.Add(1)
	entity := NewEntity(1, 100)
	m.MoveTo(entity, 50, 30)

	wait.Wait()

	// Output:
	// done
}

func ExampleTwoDimensional_StopMove() {
	m := moving2.NewTwoDimensional(moving2.WithTwoDimensionalTimeUnit[int64, float64](time.Second))
	defer func() {
		m.Release()
	}()

	var wait sync.WaitGroup
	m.RegPosition2DChangeEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64], oldX, oldY float64) {
		fmt.Println("move")
	})
	m.RegPosition2DStopMoveEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println("stop")
		wait.Done()
	})
	m.RegPosition2DDestinationEvent(func(moving *moving2.TwoDimensional[int64, float64], entity moving2.TwoDimensionalEntity[int64, float64]) {
		fmt.Println("done")
		wait.Done()
	})

	wait.Add(1)
	entity := NewEntity(1, 100)
	m.MoveTo(entity, 50, 300)
	m.StopMove(1)

	wait.Wait()

	// Output:
	// stop
}
