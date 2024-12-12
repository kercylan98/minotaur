package fiber

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/kercylan98/minotaur/engine/socket"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components"
)

var (
	_ application.Component            = (*fiberWebSocketComponent)(nil)
	_ application.ComponentImporter    = (*fiberWebSocketComponent)(nil)
	_ application.ComponentInitializer = (*fiberWebSocketComponent)(nil)
)

func NewFiberWebSocketComponent(path string, provider socket.Provider) application.Component {
	return &fiberWebSocketComponent{
		path:     path,
		provider: provider,
	}
}

type fiberWebSocketComponent struct {
	fiber         components.FiberComponent
	socketFactory socket.Factory
	path          string
	provider      socket.Provider
}

func (w *fiberWebSocketComponent) OnImport(provider *application.ComponentProvider) {
	w.fiber = application.ProvideComponent[components.FiberComponent](provider)
}

func (w *fiberWebSocketComponent) OnInitialize(app *application.Context) error {
	w.socketFactory = socket.NewFactory(app.ActorSystem())
	w.fiber.RegisterFiberHandler(w.onWebSocket)
	return nil
}

func (w *fiberWebSocketComponent) OnStart(app *application.Context) {

}

func (w *fiberWebSocketComponent) onWebSocket(app *fiber.App) {
	app.Get(w.path, websocket.New(func(conn *websocket.Conn) {
		socket.ProduceFiberSocketV2(w.socketFactory, conn, w.provider.Provide())
	}))
}
