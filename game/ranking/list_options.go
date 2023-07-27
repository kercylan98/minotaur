package ranking

import "github.com/kercylan98/minotaur/utils/generic"

type ListOption[CompetitorID comparable, Score generic.Ordered] func(list *List[CompetitorID, Score])

// WithListCount 通过限制排行榜竞争者数量来创建排行榜
//   - 默认情况下允许100位竞争者
func WithListCount[CompetitorID comparable, Score generic.Ordered](rankCount int) ListOption[CompetitorID, Score] {
	return func(list *List[CompetitorID, Score]) {
		if rankCount <= 0 {
			rankCount = 1
		}
		list.rankCount = rankCount
	}
}

// WithListASC 通过升序的方式创建排行榜
//   - 默认情况下为降序
func WithListASC[CompetitorID comparable, Score generic.Ordered]() ListOption[CompetitorID, Score] {
	return func(list *List[CompetitorID, Score]) {
		list.asc = true
	}
}
