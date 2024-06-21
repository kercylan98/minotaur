package transport

import "time"

type ServerStatus struct {
	ConnectionNum            int       `json:"connection_num"`              // 连接数
	LastConnectionOpenedTime time.Time `json:"last_connection_opened_time"` // 最后一次连接打开时间
	LastConnectionClosedTime time.Time `json:"last_connection_closed_time"` // 最后一次连接关闭时间
}
