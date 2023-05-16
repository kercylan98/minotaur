package corss

type Message struct {
	ServerId int64  `json:"server_id"`
	Packet   []byte `json:"packet"`
}
