package core_test

import (
	core2 "github.com/kercylan98/minotaur/core"
	"testing"
)

func BenchmarkProcessRegisterTable_Register(b *testing.B) {
	table := core2.NewProcessManager("", 128)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		table.Register(&TestProcess{})
	}
	b.StopTimer()
}

type TestProcess struct {
}

func (t *TestProcess) GetAddress() core2.Address {
	return ""
}

func (t *TestProcess) Deaden() bool {
	return false
}

func (t *TestProcess) Dead() {

}

func (t *TestProcess) SendUserMessage(sender *core2.ProcessRef, message core2.Message) {

}

func (t *TestProcess) SendSystemMessage(sender *core2.ProcessRef, message core2.Message) {

}

func (t *TestProcess) Terminate(ref *core2.ProcessRef) {

}
