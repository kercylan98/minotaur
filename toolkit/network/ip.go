package network

import (
	"fmt"
	"net"
	"strconv"
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

// NormalizeAddress 标准化输入地址
func NormalizeAddress(addr string) (string, int, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}

	if host == "" || host == "localhost" {
		host = "127.0.0.1"
	}

	ip := net.ParseIP(host)
	if ip == nil {
		return "", 0, fmt.Errorf("invalid host: %s", host)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return "", 0, fmt.Errorf("invalid port: %s", portStr)
	}

	return ip.String(), port, nil
}

// IsLocalAddress 检查目标地址是否为本地地址
func IsLocalAddress(targetAddr string) (bool, error) {
	normalizedTarget, targetPort, err := NormalizeAddress(targetAddr)
	if err != nil {
		return false, err
	}

	ip := net.ParseIP(normalizedTarget)
	if ip == nil {
		return false, fmt.Errorf("invalid address: %s", normalizedTarget)
	}

	// 检查本地接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return false, err
	}

	for _, ifa := range interfaces {
		addresses, err := ifa.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addresses {
			if ip.Equal(a.(*net.IPNet).IP) && targetPort == 8888 { // 替换为实际要检查的端口
				return true, nil
			}
		}
	}

	return false, nil
}

// IsSameLocalAddress 比较两个地址是否为同一本地地址
func IsSameLocalAddress(addr1, addr2 string) bool {
	ip1, port1, err := NormalizeAddress(addr1)
	if err != nil {
		return false
	}

	ip2, port2, err := NormalizeAddress(addr2)
	if err != nil {
		return false
	}

	return ip1 == ip2 && port1 == port2
}
