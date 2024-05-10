package ranking

import "github.com/kercylan98/minotaur/toolkit/constraints"

type (
	BinarySearchRankChangeEventHandle[CompetitorID comparable, Score constraints.Ordered]      func(ranking *BinarySearch[CompetitorID, Score], competitorId CompetitorID, oldRank, newRank int, oldScore, newScore Score)
	BinarySearchRankClearBeforeEventHandle[CompetitorID comparable, Score constraints.Ordered] func(ranking *BinarySearch[CompetitorID, Score])
)

type binarySearchEvent[CompetitorID comparable, Score constraints.Ordered] struct {
	rankChangeEventHandles      []BinarySearchRankChangeEventHandle[CompetitorID, Score]
	rankClearBeforeEventHandles []BinarySearchRankClearBeforeEventHandle[CompetitorID, Score]
}

// RegRankChangeEventHandle 注册排行榜变更事件
func (slf *binarySearchEvent[CompetitorID, Score]) RegRankChangeEventHandle(handle BinarySearchRankChangeEventHandle[CompetitorID, Score]) {
	slf.rankChangeEventHandles = append(slf.rankChangeEventHandles, handle)
}

// OnRankChangeEvent 触发排行榜变更事件
func (slf *binarySearchEvent[CompetitorID, Score]) OnRankChangeEvent(list *BinarySearch[CompetitorID, Score], competitorId CompetitorID, oldRank, newRank int, oldScore, newScore Score) {
	for _, handle := range slf.rankChangeEventHandles {
		handle(list, competitorId, oldRank, newRank, oldScore, newScore)
	}
}

// RegRankClearBeforeEventHandle 注册排行榜清空前事件
func (slf *binarySearchEvent[CompetitorID, Score]) RegRankClearBeforeEventHandle(handle BinarySearchRankClearBeforeEventHandle[CompetitorID, Score]) {
	slf.rankClearBeforeEventHandles = append(slf.rankClearBeforeEventHandles, handle)
}

// OnRankClearBeforeEvent 触发排行榜清空前事件
func (slf *binarySearchEvent[CompetitorID, Score]) OnRankClearBeforeEvent(list *BinarySearch[CompetitorID, Score]) {
	for _, handle := range slf.rankClearBeforeEventHandles {
		handle(list)
	}
}
