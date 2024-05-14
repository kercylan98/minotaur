package vivid

import (
	"encoding/binary"
	"fmt"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"regexp"
	"strconv"
	"strings"
)

const (
	actorIdPrefix = "minotaur"
)

// ActorId 是一个 Actor 的唯一标识符，该标识符是由紧凑的不可读字符串组成，其中包含了 Actor 完整的资源定位信息
//   - minotaur://my-system/user/my-localActor
//   - minotaur.tcp://localhost:1234/user/my-localActor
//   - minotaur.tcp://my-cluster@localhost:1234/user/my-localActor
type ActorId string

func NewActorId(network, cluster, host string, port uint16, system, name string) ActorId {
	networkLen := uint16(len(network))
	clusterLen := uint16(len(cluster))
	hostLen := uint16(len(host))
	systemLen := uint16(len(system))
	nameLen := uint16(len(name))

	// 计算需要的字节数
	size := networkLen + clusterLen + hostLen + systemLen + nameLen + 12 // 添加端口号和长度信息

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
	binary.BigEndian.PutUint16(actorId[offset:], nameLen)
	offset += 2

	// 写入网络信息
	copy(actorId[offset:], network)
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

	// 写入名称信息
	copy(actorId[offset:], name)
	offset += nameLen

	// 写入端口信息
	binary.BigEndian.PutUint16(actorId[offset:], port)

	// 转换为字符串
	return ActorId(actorId)
}

// ParseActorId 用于解析可读的 ActorId 字符串为 ActorId 对象
//   - minotaur://my-system/user/my-localActor
//   - minotaur.tcp://localhost:1234/user/my-localActor
//   - minotaur.tcp://my-cluster@localhost:1234/user/my-localActor
func ParseActorId(actorId string) (ActorId, error) {
	var network, cluster, host, system, name string
	var port int
	var portStr string

	// 定义正则表达式来匹配不同格式的 ActorId
	re1 := regexp.MustCompile(`^` + actorIdPrefix + `://([^/@]+@)?([^/:]+):(\d+)/([^/]+)/([^/]+)$`)
	re2 := regexp.MustCompile(`^` + actorIdPrefix + `\.tcp://([^/:]+):(\d+)/([^/]+)/([^/]+)$`)
	re3 := regexp.MustCompile(`^` + actorIdPrefix + `://([^/]+)/([^/]+)/([^/]+)$`)

	if matches := re1.FindStringSubmatch(actorId); matches != nil {
		cluster = matches[1]
		host = matches[2]
		portStr = matches[3]
		system = matches[4]
		name = matches[5]
		network = "tcp"
	} else if matches := re2.FindStringSubmatch(actorId); matches != nil {
		host = matches[1]
		portStr = matches[2]
		system = matches[3]
		name = matches[4]
		network = "tcp"
	} else if matches := re3.FindStringSubmatch(actorId); matches != nil {
		system = matches[1]
		name = matches[2]
		network = ""
		cluster = ""
		host = ""
		portStr = "0"
	} else {
		return "", fmt.Errorf("%w: %s", ErrActorIdInvalid, actorId)
	}

	// 将端口号从字符串转换为整数
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", fmt.Errorf("%w: %s, %w", ErrActorIdInvalid, actorId, err)
	}

	return NewActorId(network, cluster, host, uint16(port), system, name), nil
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

// Name 获取 ActorId 的名称信息
func (a ActorId) Name() string {
	networkLen := binary.BigEndian.Uint16([]byte(a[:2]))
	clusterLen := binary.BigEndian.Uint16([]byte(a[2:4]))
	hostLen := binary.BigEndian.Uint16([]byte(a[4:6]))
	systemLen := binary.BigEndian.Uint16([]byte(a[6:8]))
	nameLen := binary.BigEndian.Uint16([]byte(a[8:10]))
	v := a[10+networkLen+clusterLen+hostLen+systemLen : 10+networkLen+clusterLen+hostLen+systemLen+nameLen]
	return string(v)
}

// String 获取 ActorId 的字符串表示
func (a ActorId) String() string {
	var builder strings.Builder
	builder.WriteString(actorIdPrefix)
	if network := a.Network(); network != "" {
		builder.WriteString(".")
		builder.WriteString(network)
	}
	builder.WriteString("://")
	if cluster := a.Cluster(); cluster != "" {
		builder.WriteString(cluster)
		builder.WriteString("@")
	}
	builder.WriteString(a.Host())
	if port := a.Port(); port != 0 {
		builder.WriteString(":")
		builder.WriteString(convert.Uint16ToString(port))
	}
	builder.WriteString("/")
	builder.WriteString(a.System())
	builder.WriteString("/")
	builder.WriteString(a.Name())
	return builder.String()
}
