package user

import (
	"github.com/kercylan98/minotaur/engine/vivid"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
)

var (
	_ components.UserComponent      = (*userComponent)(nil)
	_ application.ComponentImporter = (*userComponent)(nil)
)

func NewUserComponent() application.Component {
	return &userComponent{}
}

type userComponent struct {
	app *application.Context
	ref vivid.ActorRef

	fiber    components.FiberComponent
	database components.DatabaseComponent
}

func (u *userComponent) OnImport(provider *application.ComponentProvider) {
	u.fiber = application.ProvideComponent[components.FiberComponent](provider)
	u.database = application.ProvideComponent[components.DatabaseComponent](provider)

	u.fiber.RegisterFiberHandler(u.onInitRoutes)
}

func (u *userComponent) onInitRoutes(app *fiber.App) {
	app.Post("/api/v1/ucc/register", u.onRegister)
	app.Post("/api/v1/ucc/login", u.onLogin)
	app.Post("/api/v1/ucc/refresh_token", u.onRefreshToken)
}

func (u *userComponent) OnStart(app *application.Context) {
	u.app = app
	u.ref = app.ActorSystem().ActorOfF(func() vivid.Actor {
		return newActor(u)
	})
}

func (u *userComponent) onRegister(ctx *fiber.Context) error {
	return ctx.JSON(map[string]string{
		"message": "register success",
	})
}

func (u *userComponent) onLogin(ctx *fiber.Context) error {
	return ctx.JSON(map[string]string{
		"message": "login success",
	})
}

func (u *userComponent) onRefreshToken(ctx *fiber.Context) error {
	return ctx.JSON(map[string]string{
		"message": "refresh token success",
	})
}
