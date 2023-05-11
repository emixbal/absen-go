package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

// recap-member-attendance

func RecapMemberAttendance(app *fiber.App) {
	IsAuthenticated := app.Group("/recap-member-attendance", middlewares.IsAuthenticated)

	IsAuthenticated.Get("/per-class/:class_id/:year_month", controllers.RecapMemberAttendancePerClass)
	IsAuthenticated.Get("/per-member/:member_id/:year_month", controllers.RecapMemberAttendancePerMember)
}

func RecapMemberOffsiteAssignment(app *fiber.App) {
	IsAuthenticated := app.Group("/recap-offsite-assignment", middlewares.IsAuthenticated)
	IsAuthenticated.Get("/all", controllers.OffSiteAssignmentRecapAll)
}
