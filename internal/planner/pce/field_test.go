package pce_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/planner/pce"
	"testing"
)

func TestGetFieldGolangType(t *testing.T) {
	fmt.Println(pce.GetFieldGolangType(new(pce.String)))
}
