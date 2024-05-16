package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

func TestActorId_Network(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	network := actorId.Network()
	t.Log(network)

	if network != "tcp" {
		t.Fail()
	}
}

func TestActorId_Cluster(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	cluster := actorId.Cluster()
	t.Log(cluster)

	if cluster != "my-cluster" {
		t.Fail()
	}
}

func TestActorId_Host(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	host := actorId.Host()
	t.Log(host)

	if host != "localhost" {
		t.Fail()
	}
}

func TestActorId_Port(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	port := actorId.Port()
	t.Log(port)

	if port != 1234 {
		t.Fail()
	}
}

func TestActorId_System(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	system := actorId.System()
	t.Log(system)

	if system != "my-system" {
		t.Fail()
	}
}

func TestActorId_Path(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	path := actorId.Path()
	t.Log(path)

	if path != "my-localActorRef" {
		t.Fail()
	}
}

func TestActorId_Name(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	name := actorId.Name()
	t.Log(name)

	if name != "my-localActorRef" {
		t.Fail()
	}
}

func TestActorId_String(t *testing.T) {
	actorId := vivid.NewActorId("tcp", "my-cluster", "localhost", 1234, "my-system", "my-localActorRef")

	str := actorId.String()
	t.Log(str)

	if str != "minotaur.tcp://my-cluster@localhost:1234/my-system/my-localActorRef" {
		t.Fail()

	}
}

func TestActorId_Parse(t *testing.T) {
	actorId := "minotaur.tcp://my-cluster@localhost:1234/my-system/my-localActorRef"

	parsed, err := vivid.ParseActorId(actorId)
	if err != nil {
		panic(err)
	}
	t.Log(parsed.String())

	if parsed.String() != actorId {
		t.Fail()
	}
}
