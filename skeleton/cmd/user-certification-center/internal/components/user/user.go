package user

import (
	internalcomponents "github.com/kercylan98/minotaur/skeleton/cmd/user-certification-center/internal/components"
	"github.com/kercylan98/minotaur/skeleton/pkg/application"
	"github.com/kercylan98/minotaur/skeleton/pkg/components"
	"github.com/kercylan98/minotaur/skeleton/pkg/fiber"
)

var (
	_ internalcomponents.UserComponent = (*userComponent)(nil)
	_ application.ComponentImporter    = (*userComponent)(nil)
)

func NewUserComponent() application.Component {
	return &userComponent{}
}

type userComponent struct {
	fiber components.FiberComponent
}

func (u *userComponent) OnImport(provider *application.ComponentProvider) {
	u.fiber = application.ProvideComponent[components.FiberComponent](provider)

	u.fiber.RegisterFiberHandler(u.onInitRoutes)
}

func (u *userComponent) onInitRoutes(app *fiber.App) {
	app.Post("/api/v1/ucc/register", u.onRegister)
	app.Post("/api/v1/ucc/login", u.onLogin)
	app.Post("/api/v1/ucc/refresh_token", u.onRefreshToken)
}

func (u *userComponent) OnStart(app *application.Context) {

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
