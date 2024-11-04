package phi_test

import (
	"github.com/kercylan98/minotaur/toolkit/phi"
	"testing"
	"time"
)

func TestAccrualFailureDetector(t *testing.T) {
	warnings := make(chan time.Duration)
	node, err := phi.NewAccrualFailureDetector(
		8.0,
		200,
		500*time.Millisecond,
		0,
		500*time.Millisecond,
		warnings)

	if err != nil {
		t.Fatal(err)
	}

	go func() {
		for warning := range warnings {
			t.Log("warning:", warning)
		}
	}()

	for i := 0; i < 50; i++ {
		if i < 10 || i > 20 {
			node.Heartbeat()
		}
		t.Log(node.IsAvailable(), node.Phi())
		time.Sleep(500 * time.Millisecond)
	}

}
