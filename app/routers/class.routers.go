package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Class(app *fiber.App) {
	r := app.Group("/classes", middlewares.IsAuthenticated)

	r.Get("/", controllers.ClassFetchAll)

}
