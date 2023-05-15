package server

type crossMessage struct {
	toServerId int64
	ServerId   int64          `json:"server_id"`
	Queue      CrossQueueName `json:"queue"`
	Packet     []byte         `json:"packet"`
}
