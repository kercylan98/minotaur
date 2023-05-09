package builtin

import (
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRankingList(t *testing.T) {
	convey.Convey("NewRankingList", t, func() {
		convey.So(NewRankingList[string, int](), convey.ShouldNotEqual, nil)
		convey.So(NewRankingList[string, int](WithRankingListCount[string, int](50)).rankCount, convey.ShouldEqual, 50)
		convey.So(NewRankingList[string, int](WithRankingListASC[string, int]()).asc, convey.ShouldEqual, true)
	})
}

func TestRankingList_Competitor(t *testing.T) {
	convey.Convey("TestRankingList_Competitor", t, func() {
		rankingList := NewRankingList[string, int]()
		rankingList.Competitor("angle", 63)
		rankingList.Competitor("jacky", 59)
		rankingList.Competitor("dave", 86)
		rankingList.Competitor("jacky", 42)
		convey.So(rankingList.Size(), convey.ShouldEqual, 3)
	})
	convey.Convey("TestRankingList_Competitor", t, func() {
		rankingList := NewRankingList[string, int](WithRankingListCount[string, int](2))
		rankingList.Competitor("angle", 63)
		rankingList.Competitor("jacky", 59)
		rankingList.Competitor("dave", 86)
		rankingList.Competitor("jacky", 42)
		convey.So(rankingList.Size(), convey.ShouldEqual, 2)
		one, err := rankingList.GetCompetitor(0)
		convey.So(one, convey.ShouldEqual, "dave")
		convey.So(err, convey.ShouldEqual, nil)
		two, err := rankingList.GetCompetitor(1)
		convey.So(two, convey.ShouldEqual, "angle")
		convey.So(err, convey.ShouldEqual, nil)
	})
}

func TestRankingList_CompetitorIncrease(t *testing.T) {
	convey.Convey("TestRankingList_Competitor", t, func() {
		rankingList := NewRankingList[string, int]()
		rankingList.Competitor("angle", 63)
		rankingList.Competitor("dave", 86)
		rankingList.Competitor("jacky", 42)
		rankingList.CompetitorIncrease("jacky", 100)
		convey.So(rankingList.Size(), convey.ShouldEqual, 3)
		one, err := rankingList.GetCompetitor(0)
		convey.So(one, convey.ShouldEqual, "jacky")
		convey.So(err, convey.ShouldEqual, nil)
	})
}
