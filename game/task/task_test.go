package task

import (
	"fmt"
	"testing"
)

type Player struct {
	tasks map[string][]*Task
}

type Monster struct {
}

func TestCond(t *testing.T) {
	task := NewTask(WithType("T"), WithCounter(5), WithCondition(Cond("N", 5).Cond("M", 10)))
	task.AssignConditionValueAndRefresh("N", 5)
	task.AssignConditionValueAndRefresh("M", 10)

	RegisterRefreshTaskCounterEvent[*Player](task.Type, func(taskType string, trigger *Player, count int64) {
		fmt.Println("Player", count)
		for _, t := range trigger.tasks[taskType] {
			fmt.Println(t.CurrCount, t.IncrementCounter(count).Status)
		}
	})

	RegisterRefreshTaskConditionEvent[*Player](task.Type, func(taskType string, trigger *Player, condition Condition) {
		fmt.Println("Player", condition)
		for _, t := range trigger.tasks[taskType] {
			fmt.Println(t.CurrCount, t.AssignConditionValueAndRefresh("N", 5).Status)
		}
	})

	RegisterRefreshTaskCounterEvent[*Monster](task.Type, func(taskType string, trigger *Monster, count int64) {
		fmt.Println("Monster", count)
	})

	player := &Player{
		tasks: map[string][]*Task{
			task.Type: {task},
		},
	}
	OnRefreshTaskCounterEvent(task.Type, player, 1)
	OnRefreshTaskCounterEvent(task.Type, player, 2)
	OnRefreshTaskCounterEvent(task.Type, player, 3)
	OnRefreshTaskCounterEvent(task.Type, new(Monster), 3)

}
