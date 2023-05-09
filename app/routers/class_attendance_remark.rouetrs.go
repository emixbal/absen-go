package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ClassAttendanceRemark(app *fiber.App) {
	r := app.Group("/class-attendance-ramark", middlewares.IsAuthenticated)

	r.Post("/", controllers.ClassAttendanceRemarking)
}
