package game

// Actor 表示游戏中的可以放到游戏场景中的游戏对象的基本类型
//   - 需要注意 Actor 不等于 Player
//   - 在 Minotaur 中，每个网络连接可以表示一个 Player，而每个玩家可以拥有多个 Actor
//   - Actor 并非 Player 独有，场景中也可包含各类无主的 Actor
//   - 内置实现：builtin.Actor
//   - 构建函数：builtin.NewActor
type Actor interface {
	// SetGuid 设置对象的唯一标识符
	//  - 需要注意的是该函数不应该主动执行，否则可能产生意想不到的情况
	SetGuid(guid int64)
	// GetGuid 获取对象的唯一标识符
	GetGuid() int64
}
