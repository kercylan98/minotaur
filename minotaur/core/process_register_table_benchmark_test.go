package core_test

import (
	"github.com/kercylan98/minotaur/minotaur/core"
	"testing"
)

func BenchmarkProcessRegisterTable_Register(b *testing.B) {
	table := core.NewProcessManager("", 1, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table.Register(&TestProcess{})
	}
	b.StopTimer()
}

type TestProcess struct {
}

func (t *TestProcess) GetAddress() core.Address {
	return ""
}

func (t *TestProcess) Deaden() bool {
	return false
}

func (t *TestProcess) Dead() {

}

func (t *TestProcess) SendUserMessage(sender *core.ProcessRef, message core.Message) {

}

func (t *TestProcess) SendSystemMessage(sender *core.ProcessRef, message core.Message) {

}

func (t *TestProcess) Terminate(ref *core.ProcessRef) {

}
