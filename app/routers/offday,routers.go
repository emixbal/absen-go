package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Offday(app *fiber.App) {
	r := app.Group("/offday", middlewares.IsAuthenticated)

	IsSuperadmin := r.Group("/", middlewares.IsSuperadmin)

	IsSuperadmin.Get("/", controllers.OffdatFechAll)
	IsSuperadmin.Post("/", controllers.OffdayNew)
	IsSuperadmin.Delete("/:id", controllers.OffdayHardDelete)
}
