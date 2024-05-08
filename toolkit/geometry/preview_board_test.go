package geometry_test

import (
	"github.com/kercylan98/minotaur/toolkit/geometry"
	"github.com/kercylan98/minotaur/toolkit/random"
	"testing"
	"time"
)

func TestPreviewBoard2D_Start(t *testing.T) {
	var b = geometry.NewPreviewBoard2D(60, 800, 600)

	go func() {
		for {
			time.Sleep(time.Second)
			b.Update(
				geometry.NewVector(random.Int(100, 700), random.Int(100, 500)),
				geometry.NewVector(random.Int(100, 700), random.Int(100, 500)),
				geometry.NewVector(random.Int(100, 700), random.Int(100, 500)),
				geometry.NewVector(random.Int(100, 700), random.Int(100, 500)),
				geometry.NewVector(random.Int(100, 700), random.Int(100, 500)),
				geometry.NewVector(random.Int(100, 700), random.Int(100, 500)),
			)
		}
	}()

	if err := b.Start(":8088"); err != nil {
		panic(err)
	}
}
