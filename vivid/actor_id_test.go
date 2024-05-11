package vivid_test

import (
	"github.com/kercylan98/minotaur/vivid"
	"testing"
)

func TestNewActorId(t *testing.T) {
	actorId := vivid.NewActorId("127.0.0.1", 8080, "Connection", 215)
	t.Log(actorId.Host())
	t.Log(actorId.Port())
	t.Log(actorId.SystemName())
	t.Log(actorId.Guid())
}
