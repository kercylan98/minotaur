package listings_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/utils/collection/listings"
	"testing"
)

func TestPrioritySlice_Append(t *testing.T) {
	var s = listings.NewPrioritySlice[string]()
	s.Append("name_1", 2)
	s.Append("name_2", 1)
	fmt.Println(s)
}
