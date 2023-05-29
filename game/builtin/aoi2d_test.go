package builtin

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
	"time"
)

type Ent struct {
	guid         int64
	x, y, vision float64
}

func (slf *Ent) SetGuid(guid int64) {
	slf.guid = guid
}

func (slf *Ent) GetGuid() int64 {
	return slf.guid
}

func (slf *Ent) GetPosition() (x, y float64) {
	return slf.x, slf.y
}

func (slf *Ent) GetVision() float64 {
	return slf.vision
}

func TestNewAOI2D(t *testing.T) {
	aoi := NewAOI2D(10000, 10000, 100, 100)

	start := time.Now()
	for i := 0; i < 50000; i++ {
		aoi.AddEntity(&Ent{
			guid:   int64(i),
			x:      float64(random.Int(0, 10000)),
			y:      float64(random.Int(0, 10000)),
			vision: 200,
		})
	}
	fmt.Println("添加耗时：", time.Since(start))

	//start = time.Now()
	//aoi.SetAreaSize(1000, 1000)
	//fmt.Println("重设区域大小耗时：", time.Since(start))
	start = time.Now()
	aoi.SetSize(10100, 10100)
	fmt.Println("重设大小耗时：", time.Since(start))
}
