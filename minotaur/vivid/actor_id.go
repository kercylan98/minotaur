package vivid

import (
	"encoding/binary"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"net"
	"path/filepath"
	"strings"
)

const (
	actorIdPrefix    = "minotaur"
	actorIdMinLength = 12
)

// ActorId 是一个 Actor 的唯一标识符，该标识符是由紧凑的不可读字符串组成，其中包含了 Actor 完整的资源定位信息
//   - minotaur://my-system/user/my-localActorRef
//   - minotaur://localhost:1234/user/my-localActorRef
//   - minotaur://my-node@localhost:1234/user/my-localActorRef
type ActorId string

type ActorName = string
type ActorPath = string

func NewActorId(cluster, host string, port uint16, system, actorPath ActorPath) ActorId {
	if strings.HasPrefix(actorPath, "/") {
		actorPath = actorPath[1:]
	}
	networkLen := uint16(0) // Abandoned, occupying a place
	clusterLen := uint16(len(cluster))
	hostLen := uint16(len(host))
	systemLen := uint16(len(system))
	pathLen := uint16(len(actorPath))

	// 计算需要的字节数
	size := networkLen + clusterLen + hostLen + systemLen + pathLen + 12 // 添加端口号和长度信息

	// 分配内存
	actorId := make([]byte, size)
	offset := uint16(0)

	// 提前写入所有的长度信息，确保读取时可以快速定位
	binary.BigEndian.PutUint16(actorId[offset:], networkLen)
	offset += 2
	binary.BigEndian.PutUint16(actorId[offset:], clusterLen)
	offset += 2
	binary.BigEndian.PutUint16(actorId[offset:], hostLen)
	offset += 2
	binary.BigEndian.PutUint16(actorId[offset:], systemLen)
	offset += 2
	binary.BigEndian.PutUint16(actorId[offset:], pathLen)
	offset += 2

	// 写入网络信息
	copy(actorId[offset:], "")
	offset += networkLen

	// 写入集群信息
	copy(actorId[offset:], cluster)
	offset += clusterLen

	// 写入主机信息
	copy(actorId[offset:], host)
	offset += hostLen

	// 写入系统信息
	copy(actorId[offset:], system)
	offset += systemLen

	// 写入路径信息
	copy(actorId[offset:], actorPath)
	offset += pathLen

	// 写入端口信息
	binary.BigEndian.PutUint16(actorId[offset:], port)

	// 转换为字符串
	return ActorId(actorId)
}

// Invalid 检查 ActorId 是否无效
func (a ActorId) Invalid() bool {
	if len(a) < actorIdMinLength {
		return true
	}
	networkLen := binary.BigEndian.Uint16([]byte(a[:2]))
	clusterLen := binary.BigEndian.Uint16([]byte(a[2:4]))
	hostLen := binary.BigEndian.Uint16([]byte(a[4:6]))
	systemLen := binary.BigEndian.Uint16([]byte(a[6:8]))
	nameLen := binary.BigEndian.Uint16([]byte(a[8:10]))
	totalLen := actorIdMinLength + networkLen + clusterLen + hostLen + systemLen + nameLen
	if uint16(len(a)) < totalLen {
		return true
	}

	return networkLen == 0 || hostLen == 0 || systemLen == 0 || nameLen == 0
}

// Network 获取 ActorId 的网络信息
func (a ActorId) Network() string {
	length := binary.BigEndian.Uint16([]byte(a[:2]))
	v := a[10 : 10+length]
	return string(v)
}

// Cluster 获取 ActorId 的集群信息
func (a ActorId) Cluster() string {
	networkLen := binary.BigEndian.Uint16([]byte(a[:2]))
	clusterLen := binary.BigEndian.Uint16([]byte(a[2:4]))
	v := a[10+networkLen : 10+networkLen+clusterLen]
	return string(v)
}

// Host 获取 ActorId 的主机信息
func (a ActorId) Host() string {
	networkLen := binary.BigEndian.Uint16([]byte(a[:2]))
	clusterLen := binary.BigEndian.Uint16([]byte(a[2:4]))
	hostLen := binary.BigEndian.Uint16([]byte(a[4:6]))
	v := a[10+networkLen+clusterLen : 10+networkLen+clusterLen+hostLen]
	return string(v)
}

// Port 获取 ActorId 的端口信息
func (a ActorId) Port() uint16 {
	port := a[len(a)-2:]
	return binary.BigEndian.Uint16([]byte(port))
}

// System 获取 ActorId 的系统信息
func (a ActorId) System() string {
	networkLen := binary.BigEndian.Uint16([]byte(a[:2]))
	clusterLen := binary.BigEndian.Uint16([]byte(a[2:4]))
	hostLen := binary.BigEndian.Uint16([]byte(a[4:6]))
	systemLen := binary.BigEndian.Uint16([]byte(a[6:8]))
	v := a[10+networkLen+clusterLen+hostLen : 10+networkLen+clusterLen+hostLen+systemLen]
	return string(v)
}

// Path 获取 ActorId 的路径信息
func (a ActorId) Path() ActorPath {
	networkLen := binary.BigEndian.Uint16([]byte(a[:2]))
	clusterLen := binary.BigEndian.Uint16([]byte(a[2:4]))
	hostLen := binary.BigEndian.Uint16([]byte(a[4:6]))
	systemLen := binary.BigEndian.Uint16([]byte(a[6:8]))
	nameLen := binary.BigEndian.Uint16([]byte(a[8:10]))
	v := a[10+networkLen+clusterLen+hostLen+systemLen : 10+networkLen+clusterLen+hostLen+systemLen+nameLen]
	return ActorPath(v)
}

// Name 获取 ActorId 的名称信息
func (a ActorId) Name() ActorName {
	return filepath.Base(a.Path())
}

// IsLocal 检查 ActorId 是否是本地 ActorId
func (a ActorId) IsLocal(system *ActorSystem) bool {
	if a.Cluster() != system.cluster.GetClusterName() {
		return false
	}
	if a.Host() != system.cluster.GetHost() {
		return false
	}
	if a.Port() != system.cluster.GetPort() {
		return false
	}
	return true
}

// Address 获取 ActorId 的地址信息
func (a ActorId) Address() string {
	host, port := a.Host(), a.Port()
	if port == 0 {
		return host
	}
	return net.JoinHostPort(host, convert.Uint16ToString(port))
}

// String 获取 ActorId 的字符串表示
func (a ActorId) String() string {
	if a == "" {
		return ""
	}
	var builder strings.Builder
	builder.WriteString(actorIdPrefix)
	if network := a.Network(); network != "" {
		builder.WriteString(".")
		builder.WriteString(network)
	}
	builder.WriteString("://")
	host := a.Host()
	port := a.Port()
	if cluster := a.Cluster(); cluster != "" {
		if host != "" || port != 0 {
			builder.WriteString(cluster)
			builder.WriteString("@")
		} else {
			builder.WriteString(cluster)
			builder.WriteString("@/")
		}
	}
	builder.WriteString(host)
	if port != 0 {
		builder.WriteString(":")
		builder.WriteString(convert.Uint16ToString(port))
	}
	if host != "" {
		builder.WriteString("/")
	}
	builder.WriteString(a.System())
	builder.WriteString("/")
	builder.WriteString(a.Path())
	return builder.String()
}
