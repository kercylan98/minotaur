package space

// NewRoomControllerOptions 创建房间控制器选项
func NewRoomControllerOptions() *RoomControllerOptions {
	return &RoomControllerOptions{}
}

// mergeRoomControllerOptions 合并房间控制器选项
func mergeRoomControllerOptions(options ...*RoomControllerOptions) *RoomControllerOptions {
	result := NewRoomControllerOptions()
	for _, option := range options {
		if option.maxEntityCount != nil {
			result.maxEntityCount = option.maxEntityCount
		}
	}
	return result
}

type RoomControllerOptions struct {
	maxEntityCount *int    // 房间最大实体数量
	password       *string // 房间密码
}

// WithMaxEntityCount 设置房间最大实体数量
func (rco *RoomControllerOptions) WithMaxEntityCount(maxEntityCount int) *RoomControllerOptions {
	if maxEntityCount > 0 {
		rco.maxEntityCount = &maxEntityCount
	}
	return rco
}

// WithPassword 设置房间密码
func (rco *RoomControllerOptions) WithPassword(password string) *RoomControllerOptions {
	if password != "" {
		rco.password = &password
	}
	return rco
}
