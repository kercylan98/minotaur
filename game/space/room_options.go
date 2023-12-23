package space

import "github.com/kercylan98/minotaur/utils/generic"

// NewRoomControllerOptions 创建房间控制器选项
func NewRoomControllerOptions[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]]() *RoomControllerOptions[EntityID, RoomID, Entity, Room] {
	return &RoomControllerOptions[EntityID, RoomID, Entity, Room]{}
}

// mergeRoomControllerOptions 合并房间控制器选项
func mergeRoomControllerOptions[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]](options ...*RoomControllerOptions[EntityID, RoomID, Entity, Room]) *RoomControllerOptions[EntityID, RoomID, Entity, Room] {
	result := NewRoomControllerOptions[EntityID, RoomID, Entity, Room]()
	for _, option := range options {
		if option.maxEntityCount != nil {
			result.maxEntityCount = option.maxEntityCount
		}
	}
	return result
}

type RoomControllerOptions[EntityID comparable, RoomID comparable, Entity generic.IdR[EntityID], Room generic.IdR[RoomID]] struct {
	maxEntityCount      *int                                                                       // 房间最大实体数量
	password            *string                                                                    // 房间密码
	ownerInherit        bool                                                                       // 房间所有者是否继承
	ownerInheritHandler func(controller *RoomController[EntityID, RoomID, Entity, Room]) *EntityID // 房间所有者继承处理函数
}

// WithOwnerInherit 设置房间所有者是否继承，默认为 false
//   - inherit: 是否继承，当未设置 inheritHandler 且 inherit 为 true 时，将会按照随机或根据座位号顺序继承房间所有者
//   - inheritHandler: 继承处理函数，当 inherit 为 true 时，该函数将会被调用，传入当前房间中的所有实体，返回值为新的房间所有者
func (rco *RoomControllerOptions[EntityID, RoomID, Entity, Room]) WithOwnerInherit(inherit bool, inheritHandler ...func(controller *RoomController[EntityID, RoomID, Entity, Room]) *EntityID) *RoomControllerOptions[EntityID, RoomID, Entity, Room] {
	rco.ownerInherit = inherit
	if len(inheritHandler) > 0 {
		rco.ownerInheritHandler = inheritHandler[0]
	} else if inherit {
		rco.ownerInheritHandler = func(controller *RoomController[EntityID, RoomID, Entity, Room]) *EntityID {
			if e := controller.GetFirstEmptySeatEntity(); e != nil {
				var id = e.GetId()
				return &id
			}
			if e := controller.GetRandomEntity(); e != nil {
				var id = e.GetId()
				return &id
			}
			return nil
		}
	}
	return rco
}

// WithMaxEntityCount 设置房间最大实体数量
func (rco *RoomControllerOptions[EntityID, RoomID, Entity, Room]) WithMaxEntityCount(maxEntityCount int) *RoomControllerOptions[EntityID, RoomID, Entity, Room] {
	if maxEntityCount > 0 {
		rco.maxEntityCount = &maxEntityCount
	}
	return rco
}

// WithPassword 设置房间密码
func (rco *RoomControllerOptions[EntityID, RoomID, Entity, Room]) WithPassword(password string) *RoomControllerOptions[EntityID, RoomID, Entity, Room] {
	if password != "" {
		rco.password = &password
	}
	return rco
}
