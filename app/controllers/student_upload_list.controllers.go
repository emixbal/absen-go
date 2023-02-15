package controllers

import (
	"absen-go/app/models"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func MembersUploadList(c *fiber.Ctx) error {
	payload := struct {
		ClassID string `json:"class_id" xml:"class_id" form:"class_id" validate:"required"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		log.Println("err BodyParser")
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": err,
		})
	}

	if len(payload.ClassID) < 1 {
		log.Println("err payload len < 1")
		return c.Status(400).JSON(fiber.Map{
			"message": "class_id is required",
		})
	}

	file, err := c.FormFile("members_csv")
	if err != nil {
		log.Println("err members_csv")
		log.Println(err)
		return c.Status(400).JSON(fiber.Map{
			"message": "members_csv is required",
		})
	}

	if filepath.Ext(file.Filename) != ".csv" {
		return c.Status(400).JSON(fiber.Map{
			"message": "File type is invalid. Only .csv files are allowed.",
		})
	}

	// Save file to root directory:
	if err := c.SaveFile(file, fmt.Sprintf("./files/members_files_temp/%s", payload.ClassID+".csv")); err != nil {
		log.Println("err SaveFile")
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	result := models.MembersUploadList(payload.ClassID)

	if err := os.Remove("./files/members_files_temp/" + payload.ClassID + ".csv"); err != nil {
		log.Println("err remove temp")
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": err,
		})
	}

	return c.Status(result.Status).JSON(result)
}
