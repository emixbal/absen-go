package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	"absen-go/app/routers"
	"absen-go/config"
	"absen-go/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New())

	config.InitDB()
	db := config.DB
	database.InitMigration(db)

	routers.Init(app)
	app.Listen(":" + os.Getenv("APP_PORT"))
}
