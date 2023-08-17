package server

type GP struct {
	C  string // 连接 ID
	WT int    // WebSocket 类型
	D  []byte // 数据
	T  int64  // 时间戳
}
