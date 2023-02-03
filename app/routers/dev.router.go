package routers

import (
	"absen-go/app/models"
	"absen-go/config"
	"log"
	"math/rand"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
func RandNumberBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numberBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func Dev(app *fiber.App) {
	r := app.Group("/dev")

	r.Post("/", func(c *fiber.Ctx) error {
		var student models.Student
		var students []models.Student
		var res models.Response

		db := config.GetDBInstance()

		for i := 151; i < 181; i++ {
			student.ID = uint(i)
			student.Name = RandStringBytes(6)
			student.NISN = RandNumberBytes(10)
			student.ClassID = 6

			if result := db.Create(&student); result.Error != nil {
				log.Print("error CreateAUser")
				log.Print(result.Error)

				res.Status = http.StatusInternalServerError
				res.Message = "error save new record"
				return c.Status(res.Status).JSON(res)

			}

			students = append(students, student)
		}

		res.Status = http.StatusOK
		res.Message = "ok"
		res.Data = students

		return c.Status(res.Status).JSON(res)
	})
}
