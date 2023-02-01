package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AcademicYear(app *fiber.App) {
	IsAuthenticated := app.Group("/academic-years",
		middlewares.IsAuthenticated,
	)

	IsAdmin := IsAuthenticated.Group("/",
		middlewares.IsAdmin,
	)

	IsAdmin.Get("/", controllers.AcademicYearList)
	IsAdmin.Get("/:id", controllers.AcademicYearDetail)
	IsAdmin.Put("/:id", controllers.AcademicYearUpdate)
}
