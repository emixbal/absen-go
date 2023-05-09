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
		// &models.ClassAttendance{},
		&models.ClassAttendanceMember{},
		&models.ClassAttendanceRemarkType{},
		&models.ClassAttendanceRemark{},
		&models.Offday{},
		&models.RoutineOffday{},
		&models.Remark{},

		// &models.Extracurricular{},
		// &models.ExtracurricularClass{},
		// &models.ExtracurricularAttendance{},
		// &models.ExtracurricularAttendanceStudent{},
	)

	// init seeding
	SeederRole(db)
	SeederUser(db)
	RemarkType(db)
}
