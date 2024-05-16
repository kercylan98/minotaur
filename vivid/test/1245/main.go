package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/test/common"
	"time"
)

func main() {
	srv := common.NewServer(":1245")
	system := vivid.NewActorSystem("Account", vivid.NewActorSystemOptions().
		WithAddress(srv, "127.0.0.1", 1245).
		WithClientFactory(common.NewClient),
	)
	go func() {
		if err := system.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)

	localActor, err := vivid.ActorOf[*common.UserActor](system, vivid.NewActorOptions().WithName("User1"))
	if err != nil {
		panic(err)
	}

	_ = localActor.Tell("Hello, World!")

	if err = srv.Run(); err != nil {
		panic(err)
	}

	fmt.Println("Server is shutdown.")
}
