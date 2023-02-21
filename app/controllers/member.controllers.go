package controllers

import (
	"absen-go/app/models"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func FetchAllMembers(c *fiber.Ctx) error {
	limit := 100
	offset := 0
	class_filter := []string{}
	filter_id := ""

	if c.Query("per_page") != "" {
		limit, _ = strconv.Atoi(c.Query("per_page"))
	}
	if c.Query("page") != "" {
		page, _ := strconv.Atoi(c.Query("page"))
		offset = page - 1
	}
	if c.Query("class") != "" {
		class_filter = strings.Split(c.Query("class"), ",")
	}
	if c.Query("filter_id") != "" {
		filter_id = c.Query("filter_id")
	}

	result := models.FethAllMembers(limit, offset, class_filter, filter_id)
	return c.Status(result.Status).JSON(result)
}

func MembersUpdate(c *fiber.Ctx) error {
	var member models.Member
	var res models.Response

	p := struct {
		ClassID string `json:"class_id" xml:"class_id" form:"class_id" validate:"required"`
		Code    string `json:"code" xml:"code" form:"code" validate:"required"`
		Name    string `json:"name" xml:"name" form:"name" validate:"required"`
		NIS     string `json:"nis" xml:"nis" form:"nis"`
		NISN    string `json:"nisn" xml:"nisn" form:"nisn"`
		NBM     string `json:"nbm" xml:"nbm" form:"nbm"`
	}{}

	if err := c.BodyParser(&p); err != nil {
		log.Println("err BodyParser")
		log.Println(err)

		res.Status = 400
		res.Message = err.Error()
		return c.Status(res.Status).JSON(res)
	}

	member_id := c.Params("member_id")

	if p.Name == "" || p.ClassID == "" || p.Code == "" {
		res.Status = 400
		res.Message = "class_id, name and code are required"
		return c.Status(res.Status).JSON(res)
	}

	if p.ClassID == "1" {
		if p.NBM == "" {
			res.Status = 400
			res.Message = "NBM tidak boleh kosong untuk departemen tersebut"
			return c.Status(res.Status).JSON(res)
		}

		if p.NIS != "" {
			res.Status = 400
			res.Message = "departemen tersebut tidak memiliki NIS "
			return c.Status(res.Status).JSON(res)
		}

		if p.NISN != "" {
			res.Status = 400
			res.Message = "departemen tersebut tidak memiliki NISN "
			return c.Status(res.Status).JSON(res)
		}
	} else {
		if p.NBM != "" {
			res.Status = 400
			res.Message = "departemen tersebut tidak memiliki NBM"
			return c.Status(res.Status).JSON(res)
		}
		if p.NIS == "" {
			res.Status = 400
			res.Message = "NIS tidak boleh kosong untuk departemen tersebut"
			return c.Status(res.Status).JSON(res)
		}
		if p.NISN == "" {
			res.Status = 400
			res.Message = "NISN tidak boleh kosong untuk departemen tersebut"
			return c.Status(res.Status).JSON(res)
		}
	}

	class_id_int, _ := strconv.Atoi(p.ClassID)

	member.ClassID = class_id_int
	member.Name = p.Name
	member.Code = p.Code
	member.NIS = p.NIS
	member.NISN = p.NISN
	member.NBM = p.NBM

	result := models.MemberUpdate(&member, member_id)
	return c.Status(result.Status).JSON(result)
}
