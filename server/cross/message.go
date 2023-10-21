package cross

// Message 跨服消息数据结构
type Message struct {
	ServerId string `json:"server_id"`
	Packet   []byte `json:"packet"`
}
