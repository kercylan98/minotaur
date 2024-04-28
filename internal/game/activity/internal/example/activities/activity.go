package activities

import (
	"github.com/kercylan98/minotaur/game/activity"
	"github.com/kercylan98/minotaur/game/activity/internal/example/types"
	"github.com/kercylan98/minotaur/utils/super"
	"time"
)

var (
	DemoActivity = activity.DefineEntityDataActivity[int, int, string, *types.DemoActivityData](1).InitializeEntityData(func(activityId int, entityId string, data *activity.EntityDataMeta[*types.DemoActivityData]) {
		// 模拟数据库加载
		_ = super.UnmarshalJSON([]byte(`{"last_new_day": "2021-01-01 00:00:00", "data": {"login_num": 3}}`), data)
	})
)

func init() {
	// 模拟配置加载活动
	if err := activity.LoadOrRefreshActivity(1, 1, activity.NewOptions().
		WithStartTime(time.Now().Add(time.Second*3)),
	); err != nil {
		panic(err)
	}
}
