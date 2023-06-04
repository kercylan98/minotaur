package matrix

import (
	"github.com/kercylan98/minotaur/utils/g2d"
	"github.com/kercylan98/minotaur/utils/synchronization"
	"sync"
)

func NewMatch3[ItemType comparable, Item Match3Item[ItemType]](width, height int, options ...Match3Option[ItemType, Item]) *Match3[ItemType, Item] {
	match3 := &Match3[ItemType, Item]{
		matrix:                     NewMatrix[Item](width, height),
		generators:                 map[ItemType]func() Item{},
		links:                      synchronization.NewMap[int64, map[int64]bool](),
		positions:                  map[int64][2]int{},
		notNil:                     map[int]map[int]bool{},
		matchStrategy:              map[int]func(matrix [][]Item) [][]Item{},
		generateNotMatchRetryCount: 3,
	}
	for x := 0; x < width; x++ {
		match3.notNil[x] = map[int]bool{}
	}
	for _, option := range options {
		option(match3)
	}
	if len(match3.generators) == 0 {
		panic("please use WithMatch3Generator set at least one generation strategy")
	}
	if len(match3.matchStrategy) == 0 {
		panic("please use WithMatch3Strategy set at least one match strategy")
	}
	return match3
}

// Match3 基于三消类游戏的二维矩阵
//   - 提供了适合三消类游戏的功能
type Match3[ItemType comparable, Item Match3Item[ItemType]] struct {
	matrix *Matrix[Item]

	guid      int64                                       // 成员guid
	links     *synchronization.Map[int64, map[int64]bool] // 成员类型相同且相连的链接
	positions map[int64][2]int                            // 根据成员guid记录的成员位置
	notNil    map[int]map[int]bool                        // 特定位置是否不为空

	generators                 map[ItemType]func() Item               // 成员生成器
	matchStrategy              map[int]func(matrix [][]Item) [][]Item // 匹配策略
	generateNotMatchRetryCount int                                    // 生成不匹配类型重试次数
}

// GetHeight 获取高度
func (slf *Match3[ItemType, Item]) GetHeight() int {
	return slf.matrix.h
}

// GetWidth 获取宽度
func (slf *Match3[ItemType, Item]) GetWidth() int {
	return slf.matrix.w
}

// GenerateItem 在特定位置生成特定类型的成员
func (slf *Match3[ItemType, Item]) GenerateItem(x, y int, itemType ItemType) Item {
	item := slf.generators[itemType]()
	item.SetGuid(slf.getNextGuid())
	slf.set(x, y, item)
	return item
}

// Predict 预言
func (slf *Match3[ItemType, Item]) Predict() {
	// TODO
}

// GenerateItemsByNotMatch 生成一批在特定位置不会触发任何匹配规则的成员类型
//   - 这一批成员不会被加入到矩阵中，索引与位置索引相对应
//   - 无解的策略下会导致死循环
func (slf *Match3[ItemType, Item]) GenerateItemsByNotMatch(xys ...[2]int) (result []ItemType) {
	result = make([]ItemType, 0, len(xys))
	lastIndex := len(xys) - 1
	retry := 0
	backup := NewBackup(slf)
start:
	{
		for i, xy := range xys {
			x, y := g2d.PositionArrayToXY(xy)
			var match bool
			for _, f := range slf.generators {
				slf.set(x, y, f())
				for i := 1; i <= len(slf.matchStrategy); i++ {
					if len(slf.matchStrategy[i](slf.matrix.m)) > 0 {
						match = true
						break
					}
				}
				if !match {
					break
				}
			}
			if match {
				if i == lastIndex {
					if retry < slf.generateNotMatchRetryCount {
						retry++
						result = result[:0]
						backup.Restore()
						goto start
					} else {
						panic("no solution, the matrix rule is wrong or there are matching members.")
					}
				} else {
					result = result[:0]
					backup.Restore()
					goto start
				}
			}
			result = append(result, slf.matrix.m[x][y].GetType())
		}
	}
	return
}

// GetMatch 获取二维矩阵
//   - 该矩阵为克隆的，意味着任何修改都不会影响原有内容
func (slf *Match3[ItemType, Item]) GetMatch() [][]Item {
	var (
		width  = slf.GetWidth()
		height = slf.GetHeight()
		clone  = make([][]Item[ItemType], width)
	)
	for x := 0; x < width; x++ {
		ys := make([]Item, height)
		for y := 0; y < height; y++ {
			ys[y] = slf.matrix.m[x][y].Clone().(Item)
		}
		clone[x] = ys
	}
	return clone
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
	slf.guid++
	return slf.guid
}
