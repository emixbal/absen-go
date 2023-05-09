package database

import (
	"absen-go/app/models"

	"gorm.io/gorm"
)

func RemarkType(db *gorm.DB) {
	var counter int64

	db.Model(&models.ClassAttendanceRemarkType{}).Count(&counter)

	data := []models.ClassAttendanceRemarkType{
		{
			ID:   1,
			Name: "sakit",
		},
		{
			ID:   2,
			Name: "ijin",
		},
	}

	if counter == 0 {
		db.Create(&data)
	}
}
