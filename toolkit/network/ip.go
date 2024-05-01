package network

import "net"

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
