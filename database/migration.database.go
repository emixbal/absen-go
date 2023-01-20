package database

import (
	"absen-go/app/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Teacher{},
		&models.AcademicYear{},
		&models.Semester{},
		&models.Class{},
		&models.Student{},
		&models.Extracurricular{},
		&models.ClassAttendance{},
		&models.ExtracurricularAttendance{},
	)

	// init seeding
	SeederUser(db)
}
