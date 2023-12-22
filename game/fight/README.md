# Fight

[![Go doc](https://img.shields.io/badge/go.dev-reference-brightgreen?logo=go&logoColor=white&style=flat)](https://pkg.go.dev/github.com/kercylan98/minotaur/game/fight)

## TurnBased [`回合制`]((https://pkg.go.dev/github.com/kercylan98/minotaur/game/fight#TurnBased))

### 设计思路
- 在每个回合中计算下一次行动时间间隔。然后，会根据当前行动时间间隔选择下一个行动实体。
- 当选择到下一个行动实体后，进入行动阶段。在行动阶段中，会先触发 [`TurnBasedEntitySwitchEvent`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/fight#TurnBased.RegTurnBasedEntitySwitchEvent) 事件，然后开始计时。
- 当计时结束时，如果实体还未完成行动，则会触发 [`TurnBasedEntityActionTimeoutEvent`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/fight#TurnBased.RegTurnBasedEntityActionTimeoutEvent) 事件。如果实体已经完成行动，则会触发 [`TurnBasedEntityActionFinishEvent`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/fight#TurnBased.RegTurnBasedEntityActionFinishEvent) 事件。
- 当实体完成行动后，会触发 [`TurnBasedEntityActionSubmitEvent`](https://pkg.go.dev/github.com/kercylan98/minotaur/game/fight#TurnBased.RegTurnBasedEntityActionSubmitEvent) 事件。

回合制功能的设计思路主要考虑了以下几个方面：

* 灵活性：`TurnBased` 类型提供了丰富的属性和方法，可以满足不同游戏的需要。
* 可扩展性：`TurnBased` 类型还提供了 `AddCamp`、`GetCamp`、`SetActionTimeout` 等方法，可以根据需要扩展回合制功能。
* 事件驱动：回合制功能使用事件驱动的方式来通知回合制状态变化。

