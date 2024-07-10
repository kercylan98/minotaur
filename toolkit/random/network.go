package random

import (
	"net"
	"strconv"
	"strings"
)

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

// IPv4Host 返回一个随机产生的IPv4地址
func IPv4Host() string {
	var ip = make([]byte, 4)
	for i := 0; i < 4; i++ {
		ip[i] = byte(Int(0, 255))
	}
	return net.IP(ip).String()
}

// IPv4Address 返回一个随机产生的IPv4地址和端口
func IPv4Address() string {

	return strings.Join([]string{IPv4Host(), strconv.Itoa(Port())}, ":")
}

// IPv6Host 返回一个随机产生的IPv6地址
func IPv6Host() string {
	var ip = make([]byte, 16)
	for i := 0; i < 16; i++ {
		ip[i] = byte(Int(0, 255))
	}
	return net.IP(ip).String()
}

// IPv6Address 返回一个随机产生的IPv6地址和端口
func IPv6Address() string {
	return strings.Join([]string{IPv6Host(), strconv.Itoa(Port())}, ":")
}

// MAC 返回一个随机产生的MAC地址
func MAC() string {
	var mac = make([]byte, 6)
	for i := 0; i < 6; i++ {
		mac[i] = byte(Int(0, 255))
	}
	return net.HardwareAddr(mac).String()
}
