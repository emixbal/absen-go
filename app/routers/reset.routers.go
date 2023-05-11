package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Reset(app *fiber.App) {
	r := app.Group("/reset", middlewares.IsAuthenticated)

	IsSuperadmin := r.Group("/", middlewares.IsSuperadmin)
	IsSuperadmin.Post("/attendance", controllers.ResetAttendance)
	IsSuperadmin.Post("/member", controllers.ResetMember)
}
