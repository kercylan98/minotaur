package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

func TestActorId_Network(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	network := actorId.Network()
	t.Log(network)

	if network != "tcp" {
		t.Fail()
	}
}

func TestActorId_Cluster(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	cluster := actorId.Cluster()
	t.Log(cluster)

	if cluster != "my-cluster" {
		t.Fail()
	}
}

func TestActorId_Host(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	host := actorId.Host()
	t.Log(host)

	if host != "localhost" {
		t.Fail()
	}
}

func TestActorId_Port(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	port := actorId.Port()
	t.Log(port)

	if port != 1234 {
		t.Fail()
	}
}

func TestActorId_System(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	system := actorId.System()
	t.Log(system)

	if system != "my-system" {
		t.Fail()
	}
}

func TestActorId_Name(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	name := actorId.Name()
	t.Log(name)

	if name != "my-localActor" {
		t.Fail()
	}
}

func TestActorId_String(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActor")

	str := actorId.String()
	t.Log(str)

	if str != "minotaur.tcp://my-cluster@localhost:1234/my-system/my-localActor" {
		t.Fail()

	}
}

func TestActorId_Parse(t *testing.T) {
	actorId := "minotaur.tcp://my-cluster@localhost:1234/my-system/my-localActor"

	parsed, err := vivid.ParseActorId(actorId)
	if err != nil {
		panic(err)
	}
	t.Log(parsed.String())

	if parsed.String() != actorId {
		t.Fail()
	}
}
