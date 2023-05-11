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
		var member models.Member
		var members []models.Member
		var res models.Response

		db := config.GetDBInstance()

		for i := 100001; i <= 1000000; i++ {
			member.ID = uint(i)
			member.Name = RandStringBytes(6)
			member.NISN = RandNumberBytes(10)
			member.NIS = RandNumberBytes(10)
			member.Code = RandNumberBytes(10)
			member.ClassID = 2

			if result := db.Create(&member); result.Error != nil {
				log.Print("error CreateAUser")
				log.Print(result.Error)

				res.Status = http.StatusInternalServerError
				res.Message = "error save new record"
				return c.Status(res.Status).JSON(res)

			}

			members = append(members, member)
		}

		res.Status = http.StatusOK
		res.Message = "ok"

		return c.Status(res.Status).JSON(res)
	})
}
