package dispatcher_test

import (
	"fmt"
	"github.com/kercylan98/minotaur/server/internal/dispatcher"
)

func ExampleNewManager() {
	mgr := dispatcher.NewManager[string, *TestMessage](10124*16, func(dispatcher *dispatcher.Dispatcher[string, *TestMessage], message *TestMessage) {
		// do something
	})
	mgr.BindProducer("player_001", "shunt-001")
	mgr.BindProducer("player_002", "shunt-002")
	mgr.BindProducer("player_003", "shunt-sys")
	mgr.BindProducer("player_004", "shunt-sys")
	mgr.UnBindProducer("player_001")
	mgr.UnBindProducer("player_002")
	mgr.UnBindProducer("player_003")
	mgr.UnBindProducer("player_004")
	mgr.Wait()
	fmt.Println("done")
	// Output: done
}
