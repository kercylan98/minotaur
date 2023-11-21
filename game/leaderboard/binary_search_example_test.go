package leaderboard_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/game/leaderboard"
)

func ExampleNewBinarySearch() {
	bs := leaderboard.NewBinarySearch[string, int](leaderboard.WithBinarySearchCount[string, int](10))

	fmt.Println(bs != nil)
	// Output:
	// true
}

func ExampleBinarySearch_Competitor() {
	bs := leaderboard.NewBinarySearch[string, int](leaderboard.WithBinarySearchCount[string, int](10))

	scores := []int{6131, 132, 5133, 134, 135, 136, 137, 138, 139, 140, 222, 333, 444, 555, 666}
	for i := 1; i <= 15; i++ {
		bs.Competitor(fmt.Sprintf("competitor_%2d", i), scores[i-1])
	}

	for rank, competitor := range bs.GetAllCompetitor() {
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

func ExampleBinarySearch_RemoveCompetitor() {
	bs := leaderboard.NewBinarySearch[string, int](leaderboard.WithBinarySearchCount[string, int](10))

	scores := []int{6131, 132, 5133, 134, 135, 136, 137, 138, 139, 140, 222, 333, 444, 555, 666}
	for i := 1; i <= 15; i++ {
		bs.Competitor(fmt.Sprintf("competitor_%2d", i), scores[i-1])
	}
	bs.RemoveCompetitor("competitor_ 1")
	for rank, competitor := range bs.GetAllCompetitor() {
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

func ExampleBinarySearch_GetRank() {
	bs := leaderboard.NewBinarySearch[string, int](leaderboard.WithBinarySearchCount[string, int](10))

	scores := []int{6131, 132, 5133, 134, 135, 136, 137, 138, 139, 140, 222, 333, 444, 555, 666}
	for i := 1; i <= 15; i++ {
		bs.Competitor(fmt.Sprintf("competitor_%2d", i), scores[i-1])
	}

	fmt.Println(bs.GetRank("competitor_ 1"))

	// Output:
	// 0 <nil>
}
