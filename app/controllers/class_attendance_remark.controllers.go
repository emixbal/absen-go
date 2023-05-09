package controllers

import (
	"absen-go/app/models"
	"absen-go/app/requests"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func ClassAttendanceRemarking(c *fiber.Ctx) error {
	var res models.Response
	var car models.ClassAttendanceRemark

	p := new(requests.AddClassAttendanceRemark)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Message = "Err payloads"

		return c.Status(res.Status).JSON(res)
	}

	v := validate.Struct(p)
	if !v.Validate() {
		res.Status = http.StatusBadRequest
		res.Message = v.Errors.One()

		return c.Status(res.Status).JSON(res)
	}

	t, err := time.Parse("01/02/2006", p.Date)
	if err != nil {
		log.Println("OffdayNew controller parse date error")
		log.Println(err)

		res.Status = 400
		res.Message = "parse date error, check your date format. Accepted: MM/DD/YYYY"
		return c.Status(400).JSON(res)
	}

	// check if the date is greater than today
	if t.After(time.Now()) {
		res.Status = 400
		res.Message = "date cannot be greater than today"
		return c.Status(400).JSON(res)
	}

	car.Date = t
	car.Member.Code = p.Code
	car.Text = p.Remark
	car.RemarkTypeID = p.RemarkTypeID

	result := models.ClassAttendanceRemarking(p.Code, p.RemarkTypeID, p.Remark, t)

	return c.Status(result.Status).JSON(result)
}
