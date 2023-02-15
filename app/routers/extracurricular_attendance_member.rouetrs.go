package routers

import (
	"absen-go/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ExtracurricularAttendanceMember(app *fiber.App) {
	r := app.Group("/extra-attendance-member")

	r.Post("/", controllers.AddExtracurricularAttendanceMember)
}
