package main

import (
	"fmt"
	"github.com/kercylan98/minotaur/vivid"
	"github.com/kercylan98/minotaur/vivid/test/common"
	"time"
)

func main() {
	system := vivid.NewActorSystem("Account", vivid.NewActorSystemOptions().
		WithAddress(common.NewServer(":1244"), "127.0.0.1", 1244).
		WithClientFactory(common.NewClient),
	)
	go func() {
		if err := system.Run(); err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)

	localActor, err := vivid.ActorOf[*common.UserActor](system)
	if err != nil {
		panic(err)
	}

	remoteActorId := vivid.NewActorId("tcp", "", "127.0.0.1", 1245, "Account", "/user/User1")
	fmt.Println("Find:", remoteActorId.String())
	remoteActor, err := system.GetActor().ActorId(remoteActorId).One()
	if err != nil {
		panic(err)
	}

	if err = remoteActor.Tell("Hello, World!"); err != nil {
		panic(err)
	}

	if reply, err := remoteActor.Ask(10086); err != nil {
		panic(err)
	} else {
		fmt.Println("remote reply:", reply)
	}

	if err = localActor.Tell("local: Hello, World!"); err != nil {
		panic(err)
	}

	if reply, err := localActor.Ask(9999); err != nil {
		panic(err)
	} else {
		fmt.Println("local reply:", reply)
	}

	time.Sleep(time.Minute * 10)
}
