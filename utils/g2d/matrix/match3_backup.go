package matrix

import (
	"github.com/kercylan98/minotaur/utils/g2d"
	"github.com/kercylan98/minotaur/utils/hash"
	"github.com/kercylan98/minotaur/utils/synchronization"
)

func NewBackup[ItemType comparable, Item Match3Item[ItemType]](match3 *Match3[ItemType, Item]) *Match3Backup[ItemType, Item] {
	backup := &Match3Backup[ItemType, Item]{match3: match3}

	backup.guid = match3.guid
	backup.links = synchronization.NewMap[int64, map[int64]bool]()
	match3.links.Range(func(key int64, value map[int64]bool) {
		backup.links.Set(key, hash.Copy(value))
	})
	backup.positions = map[int64][2]int{}
	for key, value := range match3.positions {
		backup.positions[key] = g2d.PositionClone(value)
	}
	backup.notNil = map[int]map[int]bool{}
	for key, values := range match3.notNil {
		var notNil = map[int]bool{}
		for key, value := range values {
			notNil[key] = value
		}
		backup.notNil[key] = notNil
	}
	return backup
}

type Match3Backup[ItemType comparable, Item Match3Item[ItemType]] struct {
	match3 *Match3[ItemType, Item]

	guid      int64                                       // 成员guid
	links     *synchronization.Map[int64, map[int64]bool] // 成员类型相同且相连的链接
	positions map[int64][2]int                            // 根据成员guid记录的成员位置
	notNil    map[int]map[int]bool                        // 特定位置是否不为空
}

// Restore 还原备份
func (slf *Match3Backup[ItemType, Item]) Restore() {
	slf.match3.guid = slf.guid
	slf.match3.links = synchronization.NewMap[int64, map[int64]bool]()
	slf.links.Range(func(key int64, value map[int64]bool) {
		slf.match3.links.Set(key, hash.Copy(value))
	})
	slf.match3.positions = map[int64][2]int{}
	for key, value := range slf.positions {
		slf.match3.positions[key] = g2d.PositionClone(value)
	}
	slf.match3.notNil = map[int]map[int]bool{}
	for key, values := range slf.notNil {
		var notNil = map[int]bool{}
		for key, value := range values {
			notNil[key] = value
		}
		slf.match3.notNil[key] = notNil
	}
}
