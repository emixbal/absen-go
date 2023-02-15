package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

// recap-member-attendance

func RecapMemberAttendance(app *fiber.App) {
	IsAuthenticated := app.Group("/recap-member-attendance", middlewares.IsAuthenticated)

	IsAuthenticated.Get("/:class_id/:year_month", controllers.RecapMemberAttendance)

}
