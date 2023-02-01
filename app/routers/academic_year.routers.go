package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AcademicYear(app *fiber.App) {
	r := app.Group("/academic-years", middlewares.IsAuthenticated, middlewares.IsAdmin)

	r.Get("/", controllers.AcademicYearList)
}
