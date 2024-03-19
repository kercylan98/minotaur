package traffickers

import (
	"fmt"
	ws "github.com/gorilla/websocket"
	"github.com/kercylan98/minotaur/server/v2"
	"github.com/panjf2000/gnet/v2"
	netHttp "net/http"
)

func WebSocket[H netHttp.Handler](handler H, binder func(handler H, upgradeHandler func(writer netHttp.ResponseWriter, request *netHttp.Request) error)) server.Trafficker {
	w := &websocket[H]{
		http:   Http(handler).(*http[H]),
		binder: binder,
		upgrader: &ws.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 4096,
			CheckOrigin: func(r *netHttp.Request) bool {
				return true
			},
		},
	}
	binder(handler, w.OnUpgrade)
	return w
}

type websocket[H netHttp.Handler] struct {
	*http[H]
	binder   func(handler H, upgradeHandler func(writer netHttp.ResponseWriter, request *netHttp.Request) error)
	upgrader *ws.Upgrader
}

func (w *websocket[H]) OnBoot(options *server.Options) error {
	return w.http.OnBoot(options)
}

func (w *websocket[H]) OnTraffic(c gnet.Conn, packet []byte) {
	w.http.OnTraffic(c, packet)
}

func (w *websocket[H]) OnUpgrade(writer netHttp.ResponseWriter, request *netHttp.Request) (err error) {
	var (
		ip   string
		conn *ws.Conn
	)

	ip = request.Header.Get("X-Real-IP")
	conn, err = w.upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	fmt.Println("opened", ip)
	go func() {
		for {
			mt, data, err := conn.ReadMessage()
			if err != nil {
				continue
			}
			conn.WriteMessage(mt, data)
		}
	}()
	return nil
}
