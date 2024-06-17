package main

import (
	"github.com/kercylan98/minotaur/minotaur"
	"github.com/kercylan98/minotaur/minotaur/transport/network"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net/http"
)

func main() {
	minotaur.NewApplication(
		minotaur.WithNetwork(network.Http(":8080", func(handler *network.HttpServe) {
			handler.HandleFunc("GET /login", func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Hello, World!"))
			})
		})),
		minotaur.WithActorSystemName("login"),
	).Launch(func(app *minotaur.Application, ctx vivid.MessageContext) {
		app.GetServer()
	})
}
