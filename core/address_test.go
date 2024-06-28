package core_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/core"
	"testing"
)

var address = core.NewAddress("tcp", "test", "localhost", 8080, "/path")
var expected = "minotaur.tcp://test@localhost:8080/path"

func TestNewAddress(t *testing.T) {
	t.Log(core.NewAddress("tcp", "test", "localhost", 8080, "/path"))
}

func TestAddress_Address(t *testing.T) {
	t.Log(address.Address())

	if address.Address() != "localhost:8080" {
		t.Fail()
	}
}

func TestAddress_Network(t *testing.T) {
	t.Log(address.Network())

	if address.Network() != "tcp" {
		t.Fail()
	}
}

func TestAddress_System(t *testing.T) {
	t.Log(address.System())

	if address.System() != "test" {
		t.Fail()
	}
}

func TestAddress_Host(t *testing.T) {
	t.Log(address.Host())

	if address.Host() != "localhost" {
		t.Fail()
	}
}

func TestAddress_Port(t *testing.T) {
	t.Log(address.Port())

	if address.Port() != 8080 {
		t.Fail()
	}
}

func TestAddress_Path(t *testing.T) {
	t.Log(address.Path())

	if address.Path() != "/path" {
		t.Fail()
	}
}

func TestAddress_String(t *testing.T) {
	t.Log(address.String())

	if address.String() != expected {
		t.Fail()
	}
}

func TestParseAddress(t *testing.T) {
	addressStr := "minotaur://mySystem@localhost/path/to/resource"
	addr, err := core.ParseAddress(addressStr)
	if err != nil {
		t.Error(err)
		return
	}
	if addr.String() != addressStr {
		t.Error("Parsed address does not match original address")
	}

	fmt.Println("Parsed Address:", addr.String())
}
