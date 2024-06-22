package core

import (
	"encoding/binary"
	"github.com/kercylan98/minotaur/toolkit/convert"
	"strings"
	"unsafe"
)

const addressPrefix = "minotaur"

// Address 是一个唯一标识符，该标识符是由紧凑的不可读字符串组成，其中包含了完整的资源定位信息
// minotaur.$network://$system@$host:$port$path
type Address string

type Path = string

func NewAddress(network string, system, host string, port uint16, path Path) Address {
	networkLen, systemLen, hostLen, pathLen := len(network), len(system), len(host), len(path)
	totalLen := 8 + networkLen + systemLen + hostLen + pathLen + 2

	buf := make([]byte, totalLen)
	ptr := unsafe.Pointer(&buf[0])

	writeLength(&ptr, networkLen)
	writeLength(&ptr, systemLen)
	writeLength(&ptr, hostLen)
	writeLength(&ptr, pathLen)
	writeString(&ptr, network)
	writeString(&ptr, system)
	writeString(&ptr, host)
	writeUint16(&ptr, port)
	writeString(&ptr, path)

	return Address(buf)
}

func (a Address) Network() string {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)

	return readString(&ptr, 0, networkLen)
}

func (a Address) System() string {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)
	systemLen := readLength(&ptr, 1)

	return readString(&ptr, networkLen, systemLen)
}

func (a Address) Host() string {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)
	systemLen := readLength(&ptr, 1)
	hostLen := readLength(&ptr, 2)

	return readString(&ptr, networkLen+systemLen, hostLen)
}

func (a Address) Port() uint16 {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)
	systemLen := readLength(&ptr, 1)
	hostLen := readLength(&ptr, 2)

	return readUint16(&ptr, networkLen+systemLen+hostLen)
}

func (a Address) Path() string {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)
	systemLen := readLength(&ptr, 1)
	hostLen := readLength(&ptr, 2)
	portLen := 2
	pathLen := readLength(&ptr, 3)

	return readString(&ptr, networkLen+systemLen+hostLen+portLen, pathLen)
}

func (a Address) Address() string {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)
	systemLen := readLength(&ptr, 1)
	hostLen := readLength(&ptr, 2)

	host := readString(&ptr, networkLen+systemLen, hostLen)
	port := readUint16(&ptr, networkLen+systemLen+hostLen)
	if port == 0 {
		return host
	}
	portStr := convert.Uint16ToString(port)
	return host + ":" + portStr
}

func (a Address) String() string {
	data := []byte(a)
	ptr := unsafe.Pointer(&data[0])

	networkLen := readLength(&ptr, 0)
	systemLen := readLength(&ptr, 1)
	hostLen := readLength(&ptr, 2)
	pathLen := readLength(&ptr, 3)

	var builder strings.Builder
	builder.WriteString(addressPrefix)
	if networkLen > 0 {
		builder.WriteString(".")
		builder.WriteString(readString(&ptr, 0, networkLen))
	}
	builder.WriteString("://")
	builder.WriteString(readString(&ptr, networkLen, systemLen))
	if hostLen > 0 {
		builder.WriteString("@")
		builder.WriteString(readString(&ptr, networkLen+systemLen, hostLen))
		port := readUint16(&ptr, networkLen+systemLen+hostLen)
		if port > 0 {
			builder.WriteString(":")
			builder.WriteString(convert.Uint16ToString(port))
		}
	}

	builder.WriteString(readString(&ptr, networkLen+systemLen+hostLen+2, pathLen))
	return builder.String()
}

func writeLength(ptr *unsafe.Pointer, length int) {
	binary.LittleEndian.PutUint16((*(*[2]byte)(unsafe.Pointer(*ptr)))[:], uint16(length))
	*ptr = unsafe.Pointer(uintptr(*ptr) + 2)
}

func writeString(ptr *unsafe.Pointer, str string) {
	copy((*(*[1 << 30]byte)(*ptr))[:len(str)], str)
	*ptr = unsafe.Pointer(uintptr(*ptr) + uintptr(len(str)))
}

func writeUint16(ptr *unsafe.Pointer, value uint16) {
	binary.LittleEndian.PutUint16((*(*[2]byte)(*ptr))[:], value)
	*ptr = unsafe.Pointer(uintptr(*ptr) + 2)
}

func readLength(ptr *unsafe.Pointer, index int) int {
	return int(binary.LittleEndian.Uint16((*(*[2]byte)(unsafe.Pointer(uintptr(*ptr) + uintptr(index*2))))[:]))
}

func readString(ptr *unsafe.Pointer, start, length int) string {
	start += 8
	return string((*(*[1 << 30]byte)(*ptr))[start : start+length])
}

func readUint16(ptr *unsafe.Pointer, start int) uint16 {
	start += 8
	return binary.LittleEndian.Uint16((*(*[2]byte)(unsafe.Pointer(uintptr(*ptr) + uintptr(start))))[:])
}
