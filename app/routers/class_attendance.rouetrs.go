package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func ClassAttendance(app *fiber.App) {
	r := app.Group("/class-attendance", middlewares.IsAuthenticated)

	r.Post("/arrive", controllers.ClassAttendanceStudentArrive)
	r.Post("/leave", controllers.ClassAttendanceStudentLeave)
}
