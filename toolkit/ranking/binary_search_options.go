package ranking

import "github.com/kercylan98/minotaur/toolkit/constraints"

type BinarySearchOption[CompetitorID comparable, Score constraints.Ordered] func(list *BinarySearch[CompetitorID, Score])

// WithBinarySearchCount 通过限制排行榜竞争者数量来创建排行榜
//   - 默认情况下允许100位竞争者
func WithBinarySearchCount[CompetitorID comparable, Score constraints.Ordered](rankCount int) BinarySearchOption[CompetitorID, Score] {
	return func(bs *BinarySearch[CompetitorID, Score]) {
		if rankCount <= 0 {
			rankCount = 1
		}
		bs.rankCount = rankCount
	}
}

// WithBinarySearchASC 通过升序的方式创建排行榜
//   - 默认情况下为降序
func WithBinarySearchASC[CompetitorID comparable, Score constraints.Ordered]() BinarySearchOption[CompetitorID, Score] {
	return func(bs *BinarySearch[CompetitorID, Score]) {
		bs.asc = true
	}
}
