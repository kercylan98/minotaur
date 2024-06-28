package vivid

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"net"
)

type ActorRef = *core.ProcessRef

// NetworkRef 基于网络地址创建 ActorRef
func NetworkRef(address string) ActorRef {
	host, port, _ := net.SplitHostPort(address)
	return core.NewProcessRef(core.NewAddress("", "", host, convert.StringToUint16(port), ""))
}
