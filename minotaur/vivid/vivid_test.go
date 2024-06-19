package vivid_test

import (
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"testing"
)

func TestActorOf(t *testing.T) {
	defer vivid.TestActorSystem.Shutdown()
	vivid.ActorOf[*vivid.IneffectiveActor](&vivid.TestActorSystem)
}
