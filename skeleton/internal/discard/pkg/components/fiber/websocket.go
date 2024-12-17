package fiber

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/kercylan98/minotaur/engine/socket"
	application2 "github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/components"
	fiber3 "github.com/kercylan98/minotaur/skeleton/internal/discard/pkg/fiber"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
)

var (
	_ application.Component            = (*fiberWebSocketComponent[any])(nil)
	_ application.ComponentImporter    = (*fiberWebSocketComponent[any])(nil)
	_ application.ComponentInitializer = (*fiberWebSocketComponent[any])(nil)
)

func NewFiberWebSocketComponent[HandleFunc any](path string, provider func(router components.RouterComponent[HandleFunc]) socket.Actor) application.Component {
	return &fiberWebSocketComponent[HandleFunc]{
		path:     path,
		provider: provider,
	}
}

type fiberWebSocketComponent[HandleFunc any] struct {
	socketFactory socket.Factory
	path          string
	provider      func(router components.RouterComponent[HandleFunc]) socket.Actor

	fiber  components.FiberComponent
	router components.RouterComponent[HandleFunc]
}

func (w *fiberWebSocketComponent[HandleFunc]) OnImport(provider *application.ComponentProvider) {
	w.fiber = application.ProvideComponent[components.FiberComponent](provider)
	w.router = application.ProvideComponent[components.RouterComponent[HandleFunc]](provider)
}

func (w *fiberWebSocketComponent[HandleFunc]) OnInitialize(app *application2.Context) error {
	w.socketFactory = socket.NewFactory(app.ActorSystem())
	w.fiber.RegisterFiberHandler(w.onWebSocket)
	return nil
}

func (w *fiberWebSocketComponent[HandleFunc]) OnStart(app *application2.Context) {

}

func (w *fiberWebSocketComponent[HandleFunc]) onWebSocket(app *fiber3.App) {
	app.Get(w.path, func(ctx *fiber3.Context) error {
		websocket.New(func(conn *websocket.Conn) {
			socket.ProduceFiberSocketV2(w.socketFactory, conn, w.provider(w.router))
		})
		return nil
	})
}
