package routers

import (
	"absen-go/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ClassAttendance(app *fiber.App) {
	r := app.Group("/class-attendance")

	r.Post("/", controllers.AddClassAttendanceStudent)
}
