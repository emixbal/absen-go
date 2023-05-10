package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func OffSiteAssignment(app *fiber.App) {
	r := app.Group("/off-site-assignment", middlewares.IsAuthenticated)

	r.Post("/departure", controllers.OffSiteAssignmentDeparture)
	r.Post("/arrive", controllers.OffSiteAssignmentArrive)
}
