package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

// recap-student-attendance

func RecapStudentAttendance(app *fiber.App) {
	IsAuthenticated := app.Group("/recap-student-attendance", middlewares.IsAuthenticated)

	IsAuthenticated.Get("/:class_id/:month", controllers.RecapStudentAttendance)

}
