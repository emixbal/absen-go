package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	r := app.Group("/auth")

	r.Post("/login", controllers.LoginWithRefrehToken)
	r.Post("/register", controllers.UserRegister)
	r.Post("/refresh", controllers.RefreshToken)
	r.Post("/new-password", middlewares.IsAuthenticated, controllers.NewPasswordSelf)

}
