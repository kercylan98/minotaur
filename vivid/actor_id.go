package vivid

import (
	"encoding/binary"
	"fmt"
)

const (
	actorMaxHostLength          = 255
	actorMaxSystemNameLength    = 128
	actorHostPrefixLength       = 1 // 用于存储Host长度
	actorPortPrefixLength       = 2 // Port占用2字节
	actorSystemNamePrefixLength = 1 // 用于存储SystemName长度
	actorGuidPrefixLength       = 8 // Guid占用8字节
	actorCommaLength            = 1 // 逗号占1字节
	actorIdMaxLength            = actorHostPrefixLength + actorMaxHostLength + actorCommaLength + actorPortPrefixLength + actorCommaLength + actorSystemNamePrefixLength + actorMaxSystemNameLength + actorCommaLength + actorGuidPrefixLength
)

type (
	// ActorId 是并发计算模型的基本执行单元 ID 表示，用于标识一个 Actor。
	//  - 该 Id 是全局（集群）唯一的，并且包含必要特征信息。
	//
	// 分区格式如下，以英文半角逗号分隔，ActorId 应该包含以下信息：
	//  - | HostLength | SystemNameLength | Host | Port | SystemName | Guid |
	ActorId string

	ActorGuid = uint64
)

// NewActorId 生成一个 ActorId，用于标识一个 Actor。
func NewActorId(host string, port uint16, systemName string, guid ActorGuid) ActorId {
	hostLength := len(host)
	systemNameLength := len(systemName)

	actorId := [actorIdMaxLength]byte{}
	cursor := 2

	// HostLength
	actorId[0] = byte(hostLength)

	// SystemNameLength
	actorId[1] = byte(systemNameLength)

	// Host
	copy(actorId[cursor:], host)
	cursor += hostLength

	// Port
	binary.BigEndian.PutUint16(actorId[cursor:], port)
	cursor += actorPortPrefixLength

	// SystemName
	copy(actorId[cursor:], systemName)
	cursor += systemNameLength

	// Guid
	binary.BigEndian.PutUint64(actorId[cursor:], guid)
	cursor += actorGuidPrefixLength

	return ActorId(actorId[:cursor])
}

// Host 返回 ActorId 的 Host 信息。
func (i ActorId) Host() string {
	hostLength := int(i[0])
	start := actorHostPrefixLength + 1
	end := hostLength + start
	return string(i[start:end])
}

// Port 返回 ActorId 的 Port 信息。
func (i ActorId) Port() uint16 {
	hostLength := int(i[0])
	start := actorHostPrefixLength + hostLength + 1
	end := start + actorPortPrefixLength
	return binary.BigEndian.Uint16([]byte(i[start:end]))
}

// SystemName 返回 ActorId 的 SystemName 信息。
func (i ActorId) SystemName() string {
	hostLength := int(i[0])
	systemNameLength := int(i[1])
	start := actorHostPrefixLength + hostLength + actorCommaLength + actorPortPrefixLength
	end := start + systemNameLength
	return string(i[start:end])
}

// Guid 返回 ActorId 的 Guid 信息。
func (i ActorId) Guid() ActorGuid {
	hostLength := int(i[0])
	systemNameLength := int(i[1])
	start := actorHostPrefixLength + hostLength + actorCommaLength + actorPortPrefixLength + systemNameLength
	end := start + actorGuidPrefixLength
	return binary.BigEndian.Uint64([]byte(i[start:end]))
}

// IsZero 返回 ActorId 是否为零值。
func (i ActorId) IsZero() bool {
	return i == ""
}

// String 返回 ActorId 的字符串表示。
func (i ActorId) String() string {
	return fmt.Sprintf("%s:%d:%s:%d", i.Host(), i.Port(), i.SystemName(), i.Guid())
}
