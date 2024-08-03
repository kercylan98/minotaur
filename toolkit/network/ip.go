package network

import (
	"net"
)

// IP 返回本机出站地址
func IP() (ip net.IP, err error) {
	var conn net.Conn
	conn, err = net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	_ = conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = localAddr.IP
	return
}

// IPv4 返回本机出站 IPv4 地址
func IPv4() (ip net.IP, err error) {
	return IP()
}

// IPv6 返回本机出站 IPv6 地址
func IPv6() (ip net.IP, err error) {
	ip, err = IP()
	if err == nil {
		ip = ip.To16()
	}
	return
}
