package aoi_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/aoi"
	"github.com/kercylan98/minotaur/utils/geometry"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
	"time"
)

type Ent struct {
	guid   int64
	pos    geometry.Point[float64]
	vision float64
}

func (slf *Ent) GetTwoDimensionalEntityID() int64 {
	return slf.guid
}

func (slf *Ent) GetPosition() geometry.Point[float64] {
	return slf.pos
}

func (slf *Ent) GetVision() float64 {
	return slf.vision
}

func TestNewTwoDimensional(t *testing.T) {
	aoiTW := aoi.NewTwoDimensional[int64, float64, *Ent](10000, 10000, 100, 100)

	start := time.Now()
	for i := 0; i < 50000; i++ {
		aoiTW.AddEntity(&Ent{
			guid:   int64(i),
			pos:    geometry.NewPoint[float64](float64(random.Int64(0, 10000)), float64(random.Int64(0, 10000))),
			vision: 200,
		})
	}
	fmt.Println("添加耗时：", time.Since(start))

	//start = time.Now()
	//aoiTW.SetAreaSize(1000, 1000)
	//fmt.Println("重设区域大小耗时：", time.Since(start))
	start = time.Now()
	aoiTW.SetSize(10100, 10100)
	fmt.Println("重设大小耗时：", time.Since(start))
}
