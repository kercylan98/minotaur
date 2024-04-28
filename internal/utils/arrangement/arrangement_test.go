package arrangement_test

import (
	"errors"
	"fmt"
	"github.com/kercylan98/minotaur/utils/arrangement"
	"testing"
)

type Player struct {
	ID int
}

func (slf *Player) GetID() int {
	return slf.ID
}

func (slf *Player) Equal(item arrangement.Item[int]) bool {
	return item.GetID() == slf.GetID()
}

type Team struct {
	ID int
}

func TestArrangement_Arrange(t *testing.T) {
	var a = arrangement.NewArrangement[int, *Team]()
	a.AddArea(&Team{ID: 1}, arrangement.WithAreaConstraint[int, *Team](func(area *arrangement.Area[int, *Team], item arrangement.Item[int]) error {
		if len(area.GetItems()) >= 2 {
			return errors.New("too many")
		}
		return nil
	}))
	a.AddArea(&Team{ID: 2}, arrangement.WithAreaConstraint[int, *Team](func(area *arrangement.Area[int, *Team], item arrangement.Item[int]) error {
		if len(area.GetItems()) >= 1 {
			return errors.New("too many")
		}
		return nil
	}))
	a.AddArea(&Team{ID: 3}, arrangement.WithAreaConstraint[int, *Team](func(area *arrangement.Area[int, *Team], item arrangement.Item[int]) error {
		if len(area.GetItems()) >= 2 {
			return errors.New("too many")
		}
		return nil
	}))
	//a.AddArea(&Team{ID: 3})
	for i := 0; i < 10; i++ {
		a.AddItem(&Player{ID: i + 1})
	}

	res, no := a.Arrange()
	for _, area := range res {
		var str = fmt.Sprintf("area %d: ", area.GetAreaInfo().ID)
		for id := range area.GetItems() {
			str += fmt.Sprintf("%d ", id)
		}
		fmt.Println(str)
	}
	var noStr = "no: "
	for _, i := range no {
		noStr += fmt.Sprintf("%d ", i.GetID())
	}
	fmt.Println(noStr)
}
