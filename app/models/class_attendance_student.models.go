package models

import (
	"time"

	"gorm.io/gorm"
)

type ClassAttendanceStudent struct {
	gorm.Model
	Datang            time.Time       `json:"datang" `
	Pulang            time.Time       `json:"pulang" `
	ClassAttendance   ClassAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassAttendanceID int             `json:"class_attendance_id"`
	Student           Student         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID         int             `json:"student_id"`
}
