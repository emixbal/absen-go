package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Offday(app *fiber.App) {
	r := app.Group("/offdays", middlewares.IsAuthenticated)

	IsSuperadmin := r.Group("/", middlewares.IsSuperadmin)
	IsSuperadmin.Post("/", controllers.OffdayNew)
}
