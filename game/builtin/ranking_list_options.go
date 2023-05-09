package builtin

import "github.com/kercylan98/minotaur/utils/generic"

type RankingListOption[CompetitorID comparable, Score generic.Ordered] func(list *RankingList[CompetitorID, Score])

// WithRankingListCount 通过限制排行榜竞争者数量来创建排行榜
//   - 默认情况下允许100位竞争者
func WithRankingListCount[CompetitorID comparable, Score generic.Ordered](rankCount int) RankingListOption[CompetitorID, Score] {
	return func(list *RankingList[CompetitorID, Score]) {
		if rankCount <= 0 {
			rankCount = 1
		}
		list.rankCount = rankCount
	}
}

// WithRankingListASC 通过升序的方式创建排行榜
//   - 默认情况下为降序
func WithRankingListASC[CompetitorID comparable, Score generic.Ordered]() RankingListOption[CompetitorID, Score] {
	return func(list *RankingList[CompetitorID, Score]) {
		list.asc = true
	}
}
