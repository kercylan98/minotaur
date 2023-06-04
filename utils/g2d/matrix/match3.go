package matrix

import (
	"github.com/kercylan98/minotaur/utils/synchronization"
	"sync"
)

func NewMatch3[ItemType comparable, Item Match3Item[ItemType]](width, height int, options ...Match3Option[ItemType, Item]) *Match3[ItemType, Item] {
	match3 := &Match3[ItemType, Item]{
		matrix:     NewMatrix[Item](width, height),
		generators: map[ItemType]func() Item{},
		links:      synchronization.NewMap[int64, map[int64]bool](),
		positions:  map[int64][2]int{},
		notNil:     map[int]map[int]bool{},
	}
	for x := 0; x < width; x++ {
		match3.notNil[x] = map[int]bool{}
	}
	for _, option := range options {
		option(match3)
	}
	return match3
}

// Match3 基于三消类游戏的二维矩阵
//   - 提供了适合三消类游戏的功能
type Match3[ItemType comparable, Item Match3Item[ItemType]] struct {
	matrix     *Matrix[Item]
	guid       int64                                       // 成员guid
	generators map[ItemType]func() Item                    // 成员生成器
	revokes    []func()                                    // 撤销记录
	links      *synchronization.Map[int64, map[int64]bool] // 成员类型相同且相连的链接
	positions  map[int64][2]int                            // 根据成员guid记录的成员位置
	notNil     map[int]map[int]bool                        // 特定位置是否不为空
}

// GetHeight 获取高度
func (slf *Match3[ItemType, Item]) GetHeight() int {
	return slf.matrix.h
}

// GetWidth 获取宽度
func (slf *Match3[ItemType, Item]) GetWidth() int {
	return slf.matrix.w
}

// Revoke 撤销特定步数
func (slf *Match3[ItemType, Item]) Revoke(step int) {
	if step <= 0 {
		return
	}
	if step > len(slf.revokes) {
		step = len(slf.revokes)
	}
	for i := 0; i < step; i++ {
		slf.revokes[i]()
	}
	slf.revokes = slf.revokes[step:]
}

// RevokeAll 撤销全部
func (slf *Match3[ItemType, Item]) RevokeAll() {
	slf.Revoke(len(slf.revokes))
}

// RevokeClear 清除所有撤销记录
func (slf *Match3[ItemType, Item]) RevokeClear() {
	slf.revokes = slf.revokes[:0]
}

// GenerateItem 在特定位置生成特定类型的成员
func (slf *Match3[ItemType, Item]) GenerateItem(x, y int, itemType ItemType) Item {
	slf.addRevoke(func() {
		item := slf.matrix.m[x][y]
		slf.set(x, y, item)
	})
	item := slf.generators[itemType]()
	item.SetGuid(slf.getNextGuid())
	slf.set(x, y, item)
	return item
}

// 设置特定位置的成员
func (slf *Match3[ItemType, Item]) set(x, y int, item Item) {
	if old := slf.matrix.m[x][y]; slf.notNil[x][y] {
		oldGuid := old.GetGuid()
		for linkGuid := range slf.links.Get(oldGuid) {
			xy := slf.positions[linkGuid]
			slf.searchNeighbour(xy[0], xy[1], synchronization.NewMap[int64, bool](), synchronization.NewMap[int64, bool]())
		}
		slf.links.Delete(oldGuid)
	}

	slf.notNil[x][y] = true
	slf.matrix.Set(x, y, item)
	slf.positions[item.GetGuid()] = [2]int{x, y}
	slf.searchNeighbour(x, y, synchronization.NewMap[int64, bool](), synchronization.NewMap[int64, bool]())
}

func (slf *Match3[ItemType, Item]) searchNeighbour(x, y int, filter *synchronization.Map[int64, bool], childrenLinks *synchronization.Map[int64, bool]) {
	var (
		item           = slf.matrix.m[x][y]
		neighboursLock sync.Mutex
		neighbours     = map[int64]bool{}
		itemType       = item.GetType()
		wait           sync.WaitGroup
		itemGuid       = item.GetGuid()
		handle         = func(x, y int) bool {
			neighbour := slf.matrix.m[x][y]
			if !slf.notNil[x][y] || neighbour.GetType() != itemType {
				return false
			}
			neighbourGuid := neighbour.GetGuid()
			neighboursLock.Lock()
			neighbours[neighbourGuid] = true
			neighboursLock.Unlock()
			childrenLinks.Set(neighbourGuid, true)
			slf.searchNeighbour(x, y, filter, childrenLinks)
			return true
		}
	)
	if filter.Get(itemGuid) {
		return
	}
	filter.Set(itemGuid, true)
	wait.Add(4)
	go func() {
		for sy := y - 1; sy >= 0; sy-- {
			if !handle(x, sy) {
				break
			}
		}
		wait.Done()
	}()
	go func() {
		for sy := y + 1; sy < slf.matrix.h; sy++ {
			if !handle(x, sy) {
				break
			}
		}
		wait.Done()
	}()
	go func() {
		for sx := x - 1; sx >= 0; sx-- {
			if !handle(sx, y) {
				break
			}
		}
		wait.Done()
	}()
	go func() {
		for sx := x + 1; sx < slf.matrix.w; sx++ {
			if !handle(sx, y) {
				break
			}
		}
		wait.Done()
	}()
	wait.Wait()
	childrenLinks.Range(func(key int64, value bool) {
		neighbours[key] = value
	})
	slf.links.Set(itemGuid, neighbours)
}

// 获取下一个guid
func (slf *Match3[ItemType, Item]) getNextGuid() int64 {
	slf.addRevoke(func() {
		slf.guid--
	})
	slf.guid++
	return slf.guid
}

// 添加撤销记录
func (slf *Match3[ItemType, Item]) addRevoke(revoke func()) {
	slf.revokes = append([]func(){revoke}, slf.revokes...)
}
