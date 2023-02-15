package database

import (
	"absen-go/app/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Teacher{},
		&models.AcademicYear{},
		&models.Semester{},
		&models.Class{},
		&models.Member{},
		&models.ClassAttendance{},
		&models.ClassAttendanceMember{},

		// &models.Extracurricular{},
		// &models.ExtracurricularClass{},
		// &models.ExtracurricularAttendance{},
		// &models.ExtracurricularAttendanceStudent{},
	)

	// init seeding
	SeederRole(db)
	SeederUser(db)
}
