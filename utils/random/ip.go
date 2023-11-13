package random

import (
	"fmt"
	"net"
)

// NetIP 返回一个随机的IP地址
func NetIP() net.IP {
	return net.IPv4(byte(Int64(0, 255)), byte(Int64(0, 255)), byte(Int64(0, 255)), byte(Int64(0, 255)))
}

// Port 返回一个随机的端口号
func Port() int {
	return Int(1, 65535)
}

// IPv4 返回一个随机产生的IPv4地址。
func IPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", Int(1, 255), Int(0, 255), Int(0, 255), Int(0, 255))
}

// IPv4Port 返回一个随机产生的IPv4地址和端口。
func IPv4Port() string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", Int(1, 255), Int(0, 255), Int(0, 255), Int(0, 255), Int(1, 65535))
}
