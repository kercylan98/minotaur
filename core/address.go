package core

import (
	"encoding/binary"
	"errors"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"net"
	"strings"
	"unsafe"
)

const addressPrefix = "minotaur"
const portLen = 2
const LengthLen = 2
const LengthTotalLen = LengthLen * 4

type Address string
type Path = string

func NewRootAddress(network string, system, host string, port uint16) Address {
	return NewAddress(network, system, host, port, "")
}

func NewAddress(network string, system, host string, port uint16, path Path) Address {
	networkLen, systemLen, hostLen, pathLen := len(network), len(system), len(host), len(path)
	totalLen := LengthTotalLen + networkLen + systemLen + hostLen + pathLen + portLen

	buf := make([]byte, totalLen)

	writeAddress(buf,
		uint16(networkLen),
		uint16(systemLen),
		uint16(hostLen),
		uint16(pathLen),
		network,
		system,
		host,
		port,
		path,
	)

	return Address(unsafe.String(unsafe.SliceData(buf), len(buf)))
}

func (a Address) Network() string {
	data := []byte(a)
	networkLen := readUint16(data, 0)
	return readString(data, LengthTotalLen, networkLen)
}

func (a Address) System() string {
	data := []byte(a)
	networkLen := readUint16(data, 0)
	systemLen := readUint16(data, LengthLen)
	return readString(data, LengthTotalLen+networkLen, systemLen)

}

func (a Address) Host() string {
	data := []byte(a)
	networkLen := readUint16(data, 0)
	systemLen := readUint16(data, LengthLen)
	hostLen := readUint16(data, LengthLen*2)
	const localhost = "127.0.0.1"
	if hostLen == 0 {
		return localhost
	}
	return readString(data, LengthTotalLen+networkLen+systemLen, hostLen)
}

func (a Address) Port() uint16 {
	data := []byte(a)
	networkLen := readUint16(data, 0)
	systemLen := readUint16(data, LengthLen)
	hostLen := readUint16(data, LengthLen*2)
	return readUint16(data, LengthTotalLen+networkLen+systemLen+hostLen)
}

func (a Address) Path() string {
	data := []byte(a)
	networkLen := readUint16(data, 0)
	systemLen := readUint16(data, LengthLen)
	hostLen := readUint16(data, LengthLen*2)
	pathLen := readUint16(data, LengthLen*3)
	return readString(data, LengthTotalLen+networkLen+systemLen+hostLen+portLen, pathLen)
}

func (a Address) Address() string {
	data := []byte(a)
	networkLen := readUint16(data, 0)
	systemLen := readUint16(data, LengthLen)
	hostLen := readUint16(data, LengthLen*2)
	host := readString(data, LengthTotalLen+networkLen+systemLen, hostLen)
	port := readUint16(data, LengthTotalLen+networkLen+systemLen+hostLen)
	return net.JoinHostPort(host, convert.Uint16ToString(port))
}

func (a Address) ParseToRoot() Address {
	return NewRootAddress(a.Network(), a.System(), a.Host(), a.Port())
}

func (a Address) IsEmpty() bool {
	return a == ""
}

func (a Address) String() string {
	if len(a) == 0 {
		return ""
	}
	data := []byte(a)
	var builder strings.Builder

	builder.WriteString(addressPrefix)
	networkLen := readUint16(data, 0)
	if networkLen > 0 {
		builder.WriteString(".")
		builder.WriteString(readString(data, LengthTotalLen, networkLen))
	}
	builder.WriteString("://")
	systemLen := readUint16(data, LengthLen)
	if systemLen > 0 {
		builder.WriteString(readString(data, LengthTotalLen+networkLen, systemLen))
		builder.WriteString("@")
	}
	hostLen := readUint16(data, LengthLen*2)
	if hostLen > 0 {
		builder.WriteString(readString(data, LengthTotalLen+networkLen+systemLen, hostLen))
		port := readUint16(data, LengthTotalLen+networkLen+systemLen+hostLen)
		if port > 0 {
			builder.WriteString(":")
			builder.WriteString(convert.Uint16ToString(port))
		}
	} else {
		builder.WriteString("127.0.0.1")
		port := readUint16(data, LengthTotalLen+networkLen+systemLen+hostLen)
		if port > 0 {
			builder.WriteString(":")
			builder.WriteString(convert.Uint16ToString(port))
		}
	}
	if pathLen := readUint16(data, LengthLen*3); pathLen > 0 {
		builder.WriteString(readString(data, LengthTotalLen+networkLen+systemLen+hostLen+portLen, pathLen))
	}
	return builder.String()
}

func writeAddress(buf []byte, vs ...any) {
	offset := 0
	for _, v := range vs {
		switch val := v.(type) {
		case string:
			copy(buf[offset:], val)
			offset += len(val)
		case uint16:
			binary.LittleEndian.PutUint16(buf[offset:], val)
			offset += 2
		default:
			panic("unsupported type")
		}
	}
}

func readString(buf []byte, offset, length uint16) string {
	return string(buf[offset : offset+length])
}

func readUint16(buf []byte, offset uint16) uint16 {
	return binary.LittleEndian.Uint16(buf[offset : offset+2])
}

func ParseAddress(addressStr string) (Address, error) {

	addressStr = strings.TrimPrefix(addressStr, addressPrefix)
	addressStr = strings.TrimPrefix(addressStr, ".")
	parts := strings.Split(addressStr, "://")
	if len(parts) != 2 {
		return "", errors.New("invalid address format")
	}

	network := parts[0]
	remainder := parts[1]

	var system, host, path string
	var port uint16
	systemHostParts := strings.Split(remainder, "@")
	if len(systemHostParts) == 2 {
		system = systemHostParts[0]
		remainder = systemHostParts[1]
	} else {
		remainder = systemHostParts[0]
	}

	hostPortPathParts := strings.SplitN(remainder, "/", 2)
	hostPort := hostPortPathParts[0]
	if len(hostPortPathParts) == 2 {
		path = "/" + hostPortPathParts[1]
	}
	if !strings.Contains(hostPort, ":") {
		hostPort += ":0"
	}

	host, portStr, err := net.SplitHostPort(hostPort)
	if err != nil {
		return "", err
	}
	port = convert.StringToUint16(portStr)

	return NewAddress(network, system, host, port, path), nil
}
