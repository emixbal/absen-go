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

func OffSiteAssignmentDeparture(c *fiber.Ctx) error {
	var res models.Response

	p := new(requests.OffSiteAssignment)
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

	result, _ := models.OffSiteAssignmentDeparture(p.Code, p.Description)
	return c.Status(result.Status).JSON(result)
}

func OffSiteAssignmentArrive(c *fiber.Ctx) error {
	var res models.Response

	p := struct {
		Code string `json:"code" xml:"code" form:"code" validate:"required"`
	}{}

	if err := c.BodyParser(&p); err != nil {
		log.Println("err BodyParser")
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	v := validate.Struct(p)
	if !v.Validate() {
		res.Status = http.StatusBadRequest
		res.Message = v.Errors.One()

		return c.Status(res.Status).JSON(res)
	}

	result, _ := models.OffSiteAssignmentArrive(p.Code)
	return c.Status(result.Status).JSON(result)
}

func OffSiteAssignmentRecapAll(c *fiber.Ctx) error {
	var res models.Response

	sort_type := "ASC"
	sort_by := "date"
	start_date_raw := ""
	end_date_raw := ""

	if c.Query("sort_type") != "" {
		if c.Query("sort_type") == "DESC" || c.Query("sort_type") == "ASC" {
			sort_type = c.Query("sort_type")
		} else {
			res.Status = 400
			res.Message = "sort_type accepted: (ASC|DESC)"
			return c.Status(400).JSON(res)
		}
	}

	if c.Query("sort_by") != "" {
		if c.Query("sort_by") == "class" || c.Query("sort_by") == "departure" || c.Query("sort_by") == "name" {
			sort_by = c.Query("sort_by")
		} else {
			res.Status = 400
			res.Message = "sort_by accepted: (class|departure|name)"
			return c.Status(400).JSON(res)
		}
	}

	if c.Query("start_date") != "" {
		start_date_raw = c.Query("start_date")
	} else {
		start_date_raw = time.Now().Format("01/02/2006")
	}

	if c.Query("end_date") != "" {
		end_date_raw = c.Query("end_date")
	} else {
		end_date_raw = time.Now().Format("01/02/2006")

	}

	start_date, err_start_date := time.Parse("01/02/2006", start_date_raw)
	if err_start_date != nil {
		log.Println("OffSiteAssignmentRecapAll controller parse date error")
		log.Println(err_start_date)

		res.Status = 400
		res.Message = "parse date error, check your date format. Accepted: MM/DD/YYYY"
		return c.Status(400).JSON(res)
	}

	end_date, err_end_date := time.Parse("01/02/2006", end_date_raw)
	if err_end_date != nil {
		log.Println("OffSiteAssignmentRecapAll controller parse date err_end_dateor")
		log.Println(err_end_date)

		res.Status = 400
		res.Message = "parse date error, check your date format. Accepted: MM/DD/YYYY"
		return c.Status(400).JSON(res)
	}

	result := models.OffSiteAssignmentRecapAll(sort_by, sort_type, start_date, end_date)

	return c.Status(200).JSON(result)
}
