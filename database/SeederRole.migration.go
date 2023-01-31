package database

import (
	"absen-go/app/models"

	"gorm.io/gorm"
)

func SeederRole(db *gorm.DB) {
	var counter int64

	db.Model(&models.Role{}).Count(&counter)

	data := []models.Role{
		{
			ID:       1,
			Name:     "superadmin",
			IsActive: true,
		},
		{
			ID:       2,
			Name:     "admin",
			IsActive: true,
		},
		{
			ID:       3,
			Name:     "coordinator",
			IsActive: true,
		},
	}

	if counter == 0 {
		db.Create(&data)
	}
}
