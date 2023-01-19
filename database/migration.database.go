package database

import (
	"absen-go/app/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		//new approach//

	)

	// init seeding
	SeederUser(db)
}
