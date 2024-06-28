package vivid

import (
	core2 "github.com/kercylan98/minotaur/core"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"net"
)

type ActorRef = *core2.ProcessRef

// NetworkRef 基于网络地址创建 ActorRef
func NetworkRef(address string) ActorRef {
	host, port, _ := net.SplitHostPort(address)
	return core2.NewProcessRef(core2.NewAddress("", "", host, convert.StringToUint16(port), ""))
}
