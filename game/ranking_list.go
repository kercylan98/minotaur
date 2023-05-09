package game

// RankingList 排行榜
type RankingList[CompetitorID comparable, Score comparable] interface {
	// Competitor 声明排行榜竞争者
	//  - 如果竞争者存在的情况下，会更新已有成绩，否则新增竞争者
	Competitor(competitorId CompetitorID, score Score)
	// RemoveCompetitor 删除特定竞争者
	RemoveCompetitor(competitorId CompetitorID)
	// GetRank 获取竞争者排名
	GetRank(competitorId CompetitorID) (int, error)
	// GetCompetitor 获取特定排名的竞争者
	GetCompetitor(rank int) (CompetitorID, error)
	// GetCompetitorWithRange 获取第start名到第end名竞争者
	GetCompetitorWithRange(start, end int) ([]CompetitorID, error)
	// GetScore 获取竞争者成绩
	GetScore(competitorId CompetitorID) (Score, error)
	// GetAllCompetitor 获取所有竞争者ID
	//  - 结果为名次有序的
	GetAllCompetitor() []CompetitorID
	// Clear 清空排行榜
	Clear()
}
