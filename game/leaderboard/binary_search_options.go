package leaderboard

import "github.com/kercylan98/minotaur/utils/generic"

type BinarySearchOption[CompetitorID comparable, Score generic.Ordered] func(list *BinarySearch[CompetitorID, Score])

// WithBinarySearchCount 通过限制排行榜竞争者数量来创建排行榜
//   - 默认情况下允许100位竞争者
func WithBinarySearchCount[CompetitorID comparable, Score generic.Ordered](rankCount int) BinarySearchOption[CompetitorID, Score] {
	return func(bs *BinarySearch[CompetitorID, Score]) {
		if rankCount <= 0 {
			rankCount = 1
		}
		bs.rankCount = rankCount
	}
}

// WithBinarySearchASC 通过升序的方式创建排行榜
//   - 默认情况下为降序
func WithBinarySearchASC[CompetitorID comparable, Score generic.Ordered]() BinarySearchOption[CompetitorID, Score] {
	return func(bs *BinarySearch[CompetitorID, Score]) {
		bs.asc = true
	}
}
