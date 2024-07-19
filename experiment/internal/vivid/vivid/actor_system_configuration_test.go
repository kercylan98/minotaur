package vivid_test

import (
	"github.com/kercylan98/minotaur/experiment/internal/vivid/vivid"
	"testing"
)

func TestActorSystemConfiguration_WithShared(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if err := recover(); err == nil {
				t.Error("expect panic, but not")
			}
		}()
		vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.WithShared(true)
		})).Shutdown(true)
	})

	t.Run("good", func(t *testing.T) {
		vivid.NewActorSystem(vivid.FunctionalActorSystemConfigurator(func(config *vivid.ActorSystemConfiguration) {
			config.
				WithPhysicalAddress(":8080").
				WithShared(true)
		})).Shutdown(true)
	})
}
