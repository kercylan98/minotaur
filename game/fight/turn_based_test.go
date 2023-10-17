package fight_test

import (
	"github.com/kercylan98/minotaur/game/fight"
	"testing"
	"time"
)

type Camp struct {
	id string
}

func (slf *Camp) GetId() string {
	return slf.id
}

type Entity struct {
	id    string
	speed float64
}

func (slf *Entity) GetId() string {
	return slf.id
}

func TestTurnBased_Run(t *testing.T) {
	tbi := fight.NewTurnBased[string, string, *Camp, *Entity](func(camp *Camp, entity *Entity) time.Duration {
		return time.Duration(float64(time.Second) / entity.speed)
	})

	tbi.SetActionTimeout(func(camp *Camp, entity *Entity) time.Duration {
		return time.Second * 5
	})

	tbi.RegTurnBasedEntityActionTimeoutEvent(func(controller fight.TurnBasedControllerInfo[string, string, *Camp, *Entity]) {
		t.Log("时间", time.Now().Unix(), "回合", controller.GetRound(), "阵营", controller.GetCamp().GetId(), "实体", controller.GetEntity().GetId(), "超时")
	})

	tbi.RegTurnBasedEntitySwitchEvent(func(controller fight.TurnBasedControllerAction[string, string, *Camp, *Entity]) {
		switch controller.GetEntity().GetId() {
		case "1":
			go func() {
				time.Sleep(time.Second * 2)
				controller.Finish()
			}()
		case "2":
			controller.Stop()
		}
		t.Log("时间", time.Now().Unix(), "回合", controller.GetRound(), "阵营", controller.GetCamp().GetId(), "实体", controller.GetEntity().GetId(), "开始行动")
	})

	tbi.AddCamp(&Camp{id: "1"}, &Entity{id: "1", speed: 1}, &Entity{id: "2", speed: 1})
	tbi.AddCamp(&Camp{id: "2"}, &Entity{id: "3", speed: 1}, &Entity{id: "4", speed: 1})

	tbi.Run()
}
