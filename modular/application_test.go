package modular_test

import (
	"github.com/kercylan98/minotaur/modular"
	"testing"
)

func TestNewApplication(t *testing.T) {
	if modular.NewApplication() == nil {
		t.Fail()
	}
}

func TestRun(t *testing.T) {
	application := modular.NewApplication()
	modular.RegisterService[*AccountService, AccountServiceExposer](application)
	modular.RegisterService[*ConfigService, ConfigServiceExposer](application)
	modular.Run(application)
}
