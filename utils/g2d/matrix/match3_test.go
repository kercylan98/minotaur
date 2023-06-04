package matrix

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/g2d"
	"github.com/kercylan98/minotaur/utils/random"
	"testing"
	"time"
)

type Item[Type comparable] struct {
	guid int64
	t    Type
}

func (slf *Item[Type]) SetGuid(guid int64) {
	slf.guid = guid
}

func (slf *Item[Type]) GetGuid() int64 {
	return slf.guid
}

func (slf *Item[Type]) GetType() Type {
	return slf.t
}

func (slf *Item[Type]) Clone() Match3Item[Type] {
	return &Item[Type]{
		guid: slf.guid,
		t:    slf.t,
	}
}

func TestMatch3(t *testing.T) {
	var options []Match3Option[int, *Item[int]]
	for i := 0; i < 7; i++ {
		t := i + 1
		options = append(options, WithMatch3Generator[int, *Item[int]](t, func() *Item[int] {
			return &Item[int]{t: t}
		}))
	}
	var match3 = NewMatch3[int, *Item[int]](3, 3,
		options...,
	)

	for x := 0; x < match3.GetWidth(); x++ {
		for y := 0; y < match3.GetHeight(); y++ {
			match3.GenerateItem(x, y, random.Int(1, 2))
		}
	}

	for y := 0; y < match3.GetHeight(); y++ {
		for x := 0; x < match3.GetWidth(); x++ {
			fmt.Print(match3.matrix.m[x][y].t, " ")
		}
		fmt.Println()
	}
	fmt.Println()
	links := match3.links.Get(4)
	linkItem := match3.matrix.m[match3.positions[4][0]][match3.positions[4][1]]
	fmt.Println("LINK", linkItem.t, match3.positions[4])

	for y := 0; y < match3.GetHeight(); y++ {
		for x := 0; x < match3.GetWidth(); x++ {
			item := match3.matrix.m[x][y]
			if links[item.guid] {
				fmt.Print("*", " ")
			} else {
				fmt.Print(match3.matrix.m[x][y].t, " ")
			}
		}
		fmt.Println()
	}

	var now = time.Now()
	var xys [][2]int
	for guid := range links {
		xys = append(xys, match3.positions[guid])
	}

	for _, rect := range g2d.SearchNotRepeatFullRectangle(2, 2, xys...) {
		fmt.Println(fmt.Sprintf("找到矩形: TopLeft: (%d, %d), BottomRight: (%d, %d)", rect[0][0], rect[0][1], rect[1][0], rect[1][1]))
	}
	fmt.Println("耗时", time.Since(now))

	now = time.Now()
	for _, rect := range g2d.SearchNotRepeatCross(xys...) {
		fmt.Print("找到十字：")
		for _, points := range rect {
			fmt.Print(fmt.Sprintf("{%d, %d}", points[0], points[1]))
		}
		fmt.Println()
	}
	fmt.Println("耗时", time.Since(now))

	now = time.Now()
	for _, rect := range g2d.SearchNotRepeatRightAngle(4, xys...) {
		fmt.Print("找到L形：")
		for _, points := range rect {
			fmt.Print(fmt.Sprintf("{%d, %d}", points[0], points[1]))
		}
		fmt.Println()
	}
	fmt.Println("耗时", time.Since(now))

	now = time.Now()
	for _, rect := range g2d.SearchNotRepeatT(4, xys...) {
		fmt.Print("找到T形：")
		for _, points := range rect {
			fmt.Print(fmt.Sprintf("{%d, %d}", points[0], points[1]))
		}
		fmt.Println()
	}
	fmt.Println("耗时", time.Since(now))

	now = time.Now()
	for _, rect := range g2d.SearchNotRepeatStraightLine(3, xys...) {
		fmt.Print("找到直线：")
		for _, points := range rect {
			fmt.Print(fmt.Sprintf("{%d, %d}", points[0], points[1]))
		}
		fmt.Println()
	}
	fmt.Println("耗时", time.Since(now))
}
