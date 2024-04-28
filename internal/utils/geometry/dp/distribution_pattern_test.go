package dp

import (
	"fmt"
	"testing"
)

func TestNewDistributionPattern(t *testing.T) {

	dp := NewDistributionPattern[int](func(itemA, itemB int) bool {
		return itemA == itemB
	})

	matrix := []int{1, 1, 2, 2, 2, 2, 1, 2, 2}
	dp.LoadMatrixWithPos(3, matrix)

	for pos, link := range dp.links {
		fmt.Println(pos, link, fmt.Sprintf("%p", link))
	}

	fmt.Println()

	matrix[6] = 2
	dp.Refresh(6)
	for pos, link := range dp.links {
		fmt.Println(pos, link, fmt.Sprintf("%p", link))
	}
}
