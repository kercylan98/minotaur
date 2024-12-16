package mapped

import "time"

type UserMapped struct {
	Id        string    // 用户 ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
