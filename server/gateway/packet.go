package gateway

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
)

var packetIdentifier = []byte{0xDE, 0xAD, 0xBE, 0xEF}

// MarshalGatewayOutPacket 将数据包转换为网关出网数据包
//   - | identifier(4) | ipv4(4) | port(2) | packet |
func MarshalGatewayOutPacket(addr string, packet []byte) ([]byte, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	ipBytes := net.ParseIP(host).To4()
	if ipBytes == nil {
		return nil, errors.New("invalid IPv4 address")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 0 || port > 65535 {
		return nil, errors.New("invalid port number")
	}
	portBytes := []byte{byte(port >> 8), byte(port & 0xFF)}

	result := append(packetIdentifier, ipBytes...)
	result = append(result, portBytes...)
	result = append(result, packet...)

	return result, nil
}

// UnmarshalGatewayOutPacket 将网关出网数据包转换为数据包
//   - | identifier(4) | ipv4(4) | port(2) | packet |
func UnmarshalGatewayOutPacket(data []byte) (addr string, packet []byte, err error) {
	if len(data) < 10 {
		err = errors.New("data is too short to contain an IPv4 address and a port")
		return
	}
	if !compareBytes(data[:4], packetIdentifier) {
		err = errors.New("invalid identifier")
		return
	}
	ipAddr := net.IP(data[4:8]).String()
	port := uint16(data[8])<<8 | uint16(data[9])
	addr = fmt.Sprintf("%s:%d", ipAddr, port)
	packet = data[10:]

	return addr, packet, nil
}

// MarshalGatewayInPacket 将数据包转换为网关入网数据包
//   - | ipv4(4) | port(2) | cost(4) | packet |
func MarshalGatewayInPacket(addr string, currentTime int64, packet []byte) ([]byte, error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	ipBytes := net.ParseIP(host).To4()
	if ipBytes == nil {
		return nil, errors.New("invalid IPv4 address")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 0 || port > 65535 {
		return nil, errors.New("invalid port number")
	}
	portBytes := []byte{byte(port >> 8), byte(port & 0xFF)}
	costBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(costBytes, uint32(currentTime))

	result := append(ipBytes, portBytes...)
	result = append(result, costBytes...)
	result = append(result, packet...)

	return result, nil
}

// UnmarshalGatewayInPacket 将网关入网数据包转换为数据包
//   - | ipv4(4) | port(2) | cost(4) | packet |
func UnmarshalGatewayInPacket(data []byte) (addr string, sendTime int64, packet []byte, err error) {
	if len(data) < 10 {
		err = errors.New("data is too short")
		return
	}
	ipAddr := net.IP(data[:4]).String()
	port := uint16(data[4])<<8 | uint16(data[5])
	addr = fmt.Sprintf("%s:%d", ipAddr, port)
	sendTime = int64(binary.BigEndian.Uint32(data[6:10]))
	packet = data[10:]

	return addr, sendTime, packet, nil
}

func compareBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
