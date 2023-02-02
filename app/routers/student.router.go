package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Student(app *fiber.App) {
	IsAuthenticated := app.Group("/students", middlewares.IsAuthenticated)

	IsAuthenticated.Get("/", controllers.FetchAllStudents)
}
