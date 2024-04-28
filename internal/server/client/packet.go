package client

type Packet struct {
	wst      int             // websocket 的数据类型
	data     []byte          // 数据包
	callback func(err error) // 回调函数
}
