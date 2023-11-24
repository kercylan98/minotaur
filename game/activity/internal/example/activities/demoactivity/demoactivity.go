package demoactivity

import (
	"github.com/kercylan98/minotaur/game/activity"
	"github.com/kercylan98/minotaur/game/activity/internal/example/activities"
	"github.com/kercylan98/minotaur/utils/log"
)

func init() {
	activity.RegStartedEvent(1, onActivityStart)
}

func onActivityStart(id int) {
	log.Info("activity start", log.Int("id", id), log.Any("entity", activities.DemoActivity.GetEntityData(id, "demo_entity")))
}
