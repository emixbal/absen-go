package database

import (
	"absen-go/app/models"

	"gorm.io/gorm"
)

func SeederUser(db *gorm.DB) {
	var counter int64

	db.Model(&models.User{}).Count(&counter)

	user := []models.User{
		{
			Name:     "muhammad iqbal",
			Email:    "emixbal@gmail.com",
			Password: "$2a$10$xO0eiq3.64vo1gR1cKkEE.hwn0OvafrzVI0HhsZWeb9UuXsl7bZrq", //aaaaaaaa
			RoleID:   1,
		},
	}

	if counter == 0 {
		db.Create(&user)
	}
}
