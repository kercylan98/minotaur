package ecs_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/ecs"
	"testing"
)

func TestEntityId_Generation(t *testing.T) {
	t.Log(ecs.ExportNewEntityId(1, 1))
	t.Log(uint32(ecs.ExportNewEntityId(2, 1)))
}
