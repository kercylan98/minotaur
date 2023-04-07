package game

// Conn 用户连接抽象
//
// 连接支持使用 tag 进行参数读取
type Conn interface {
	GetConn() any
}
