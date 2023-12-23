package writeloop_test

import (
	"github.com/kercylan98/minotaur/server/writeloop"
	"github.com/kercylan98/minotaur/utils/concurrent"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Message struct {
	ID int
}

var wp = concurrent.NewPool(func() *Message {
	return &Message{}
}, func(data *Message) {
	data.ID = 0
})

func TestNewWriteLoop(t *testing.T) {
	wl := writeloop.NewWriteLoop(wp, func(message *Message) error {
		t.Log(message.ID)
		return nil
	}, func(err any) {
		t.Log(err)
	})
	assert.NotNil(t, wl)
	wl.Close()
}

func TestWriteLoop_Put(t *testing.T) {
	wl := writeloop.NewWriteLoop(wp, func(message *Message) error {
		t.Log(message.ID)
		return nil
	}, func(err any) {
		t.Log(err)
	})
	assert.NotNil(t, wl)

	for i := 0; i < 100; i++ {
		m := wp.Get()
		m.ID = i
		wl.Put(m)
	}

	wl.Close()
}

func TestWriteLoop_Close(t *testing.T) {
	wl := writeloop.NewWriteLoop(wp, func(message *Message) error {
		t.Log(message.ID)
		return nil
	}, func(err any) {
		t.Log(err)
	})
	assert.NotNil(t, wl)

	for i := 0; i < 100; i++ {
		m := wp.Get()
		m.ID = i
		wl.Put(m)
	}

	wl.Close()
}
