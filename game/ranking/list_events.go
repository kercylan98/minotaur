package ranking

import "github.com/kercylan98/minotaur/utils/generic"

type (
	RankChangeEventHandle[CompetitorID comparable, Score generic.Ordered]      func(list *List[CompetitorID, Score], competitorId CompetitorID, oldRank, newRank int, oldScore, newScore Score)
	RankClearBeforeEventHandle[CompetitorID comparable, Score generic.Ordered] func(list *List[CompetitorID, Score])
)

type event[CompetitorID comparable, Score generic.Ordered] struct {
	rankChangeEventHandles      []RankChangeEventHandle[CompetitorID, Score]
	rankClearBeforeEventHandles []RankClearBeforeEventHandle[CompetitorID, Score]
}

// RegRankChangeEventHandle 注册排行榜变更事件
func (slf *event[CompetitorID, Score]) RegRankChangeEventHandle(handle RankChangeEventHandle[CompetitorID, Score]) {
	slf.rankChangeEventHandles = append(slf.rankChangeEventHandles, handle)
}

// OnRankChangeEvent 触发排行榜变更事件
func (slf *event[CompetitorID, Score]) OnRankChangeEvent(list *List[CompetitorID, Score], competitorId CompetitorID, oldRank, newRank int, oldScore, newScore Score) {
	for _, handle := range slf.rankChangeEventHandles {
		handle(list, competitorId, oldRank, newRank, oldScore, newScore)
	}
}

// RegRankClearBeforeEventHandle 注册排行榜清空前事件
func (slf *event[CompetitorID, Score]) RegRankClearBeforeEventHandle(handle RankClearBeforeEventHandle[CompetitorID, Score]) {
	slf.rankClearBeforeEventHandles = append(slf.rankClearBeforeEventHandles, handle)
}

// OnRankClearBeforeEvent 触发排行榜清空前事件
func (slf *event[CompetitorID, Score]) OnRankClearBeforeEvent(list *List[CompetitorID, Score]) {
	for _, handle := range slf.rankClearBeforeEventHandles {
		handle(list)
	}
}
