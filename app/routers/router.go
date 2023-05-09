package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(app *fiber.App) {
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("/ path")
	})

	//palece new routers below
	User(app)
	Auth(app)
	ClassAttendance(app)
	ClassAttendanceRemark(app)
	AcademicYear(app)
	Member(app)
	RecapMemberAttendance(app)
	Class(app)
	Reset(app)
	Offday(app)

	Dev(app)
}
