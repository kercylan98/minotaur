package main

import (
	"github.com/kercylan98/minotaur/example/internal/login/mods/account"
	"github.com/kercylan98/minotaur/example/internal/login/mods/router"
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
)

func main() {
	minotaur.NewApplication(
		minotaur.WithActorSystemName("login"),
		minotaur.WithNetwork(network.Http(":8080", func(ctx vivid.ActorContext, handler *network.HttpServe) {
			ctx.LoadMod(vivid.ModOf[router.Router](&router.Mod{Handler: handler}))
			ctx.LoadMod(vivid.ModOf[account.Account](&account.Mod{}))
			ctx.ApplyMod()
		})),
	).Launch()
}
