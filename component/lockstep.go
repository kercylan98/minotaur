package component

type Lockstep[ClientID comparable, Command any] interface {
	// JoinClient 加入客户端
	JoinClient(client LockstepClient[ClientID])
	// LeaveClient 离开客户端
	LeaveClient(clientId ClientID)
	// StartBroadcast 开始广播
	StartBroadcast()
	// AddCommand 增加指令
	AddCommand(command Command)
}
