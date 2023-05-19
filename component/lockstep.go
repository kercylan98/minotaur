package component

type Lockstep[ClientID comparable, Command any] interface {
	// JoinClient 加入客户端
	JoinClient(client LockstepClient[ClientID])
	// JoinClientWithFrame 加入客户端，并且指定从特定帧开始
	JoinClientWithFrame(client LockstepClient[ClientID], frameIndex int)
	// LeaveClient 离开客户端
	LeaveClient(clientId ClientID)
	// StartBroadcast 开始广播
	StartBroadcast()
	// Stop 停止广播
	Stop()
	// AddCommand 增加指令
	AddCommand(command Command)
}
