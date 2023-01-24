package routers

import (
	"absen-go/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ClassAttendanceStudent(app *fiber.App) {
	r := app.Group("/class-attendance-student")

	r.Post("/", controllers.AddClassAttendanceStudent)
}
