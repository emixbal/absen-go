package routers

import (
	"absen-go/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ExtracurricularAttendanceStudent(app *fiber.App) {
	r := app.Group("/extra-attendance-student")

	r.Post("/", controllers.AddExtracurricularAttendanceStudent)
}
