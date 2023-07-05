package builtin_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/builtin"
)

func ExampleNewRankingList() {
	ranklingList := builtin.NewRankingList[string, int](builtin.WithRankingListCount[string, int](10))

	fmt.Println(ranklingList != nil)
	// Output:
	// true
}

func ExampleRankingList_Competitor() {
	ranklingList := builtin.NewRankingList[string, int](builtin.WithRankingListCount[string, int](10))

	scores := []int{6131, 132, 5133, 134, 135, 136, 137, 138, 139, 140, 222, 333, 444, 555, 666}
	for i := 1; i <= 15; i++ {
		ranklingList.Competitor(fmt.Sprintf("competitor_%2d", i), scores[i-1])
	}

	for rank, competitor := range ranklingList.GetAllCompetitor() {
		fmt.Println(rank, competitor)
	}

	// Output:
	// 0 competitor_ 1
	// 1 competitor_ 3
	// 2 competitor_15
	// 3 competitor_14
	// 4 competitor_13
	// 5 competitor_12
	// 6 competitor_11
	// 7 competitor_10
	// 8 competitor_ 9
	// 9 competitor_ 8
}

func ExampleRankingList_RemoveCompetitor() {
	ranklingList := builtin.NewRankingList[string, int](builtin.WithRankingListCount[string, int](10))

	scores := []int{6131, 132, 5133, 134, 135, 136, 137, 138, 139, 140, 222, 333, 444, 555, 666}
	for i := 1; i <= 15; i++ {
		ranklingList.Competitor(fmt.Sprintf("competitor_%2d", i), scores[i-1])
	}
	ranklingList.RemoveCompetitor("competitor_ 1")
	for rank, competitor := range ranklingList.GetAllCompetitor() {
		fmt.Println(rank, competitor)
	}

	// Output:
	// 0 competitor_ 3
	// 1 competitor_15
	// 2 competitor_14
	// 3 competitor_13
	// 4 competitor_12
	// 5 competitor_11
	// 6 competitor_10
	// 7 competitor_ 9
	// 8 competitor_ 8
}

func ExampleRankingList_GetRank() {
	ranklingList := builtin.NewRankingList[string, int](builtin.WithRankingListCount[string, int](10))

	scores := []int{6131, 132, 5133, 134, 135, 136, 137, 138, 139, 140, 222, 333, 444, 555, 666}
	for i := 1; i <= 15; i++ {
		ranklingList.Competitor(fmt.Sprintf("competitor_%2d", i), scores[i-1])
	}

	fmt.Println(ranklingList.GetRank("competitor_ 1"))

	// Output:
	// 0 <nil>
}
