package account

import (
	"github.com/kercylan98/minotaur/example/internal/login/mods/router"
	"github.com/kercylan98/minotaur/example/internal/login/types"
	"github.com/kercylan98/minotaur/minotaur/vivid"
	"net/http"
)

type Account interface {
	vivid.Mod

	// GetAccountInfo 获取账户信息
	GetAccountInfo(username string) (*types.Account, error)
}

type Mod struct {
	router router.Router
}

func (m *Mod) OnLifeCycle(ctx vivid.ActorContext, lifeCycle vivid.ModLifeCycle) {
	switch lifeCycle {
	case vivid.ModLifeCycleOnPreload:
		m.router = vivid.InvokeMod[router.Router](ctx)
	case vivid.ModLifeCycleOnStart:
		m.router.HandleFunc("GET /login", m.onLogin)
	default:
	}
}

func (m *Mod) GetAccountInfo(username string) (*types.Account, error) {
	return &types.Account{
		Username: username,
		Password: "123456",
	}, nil
}

func (m *Mod) onLogin(writer http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")

	account, err := m.GetAccountInfo(username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte(err.Error()))
		return
	}

	if account.Password != password {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("password error"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("login success"))
}
