package client_test

import (
	"github.com/kercylan98/minotaur/internal/server/client"
	"testing"
)

func TestClient_WriteWS(t *testing.T) {
	client.NewWebsocket("ws://127.0.0.1:9999")
}
