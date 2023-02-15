package routers

import (
	"absen-go/app/controllers"
	"absen-go/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Member(app *fiber.App) {
	IsAuthenticated := app.Group("/members", middlewares.IsAuthenticated)

	IsAuthenticated.Get("/", controllers.FetchAllMembers)
	IsAuthenticated.Post("/", controllers.MembersUploadList)
}
