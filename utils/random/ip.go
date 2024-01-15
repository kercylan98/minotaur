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

// UsablePort 随机返回一个可用的端口号，如果没有可用端口号则返回 -1
func UsablePort() int {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:0")
	if err != nil {
		return -1
	}
	cli, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return -1
	}
	defer func() { _ = cli.Close() }()
	return cli.Addr().(*net.TCPAddr).Port
}

// IPv4 返回一个随机产生的IPv4地址。
func IPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d", Int(1, 255), Int(0, 255), Int(0, 255), Int(0, 255))
}

// IPv4Port 返回一个随机产生的IPv4地址和端口。
func IPv4Port() string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", Int(1, 255), Int(0, 255), Int(0, 255), Int(0, 255), Int(1, 65535))
}
