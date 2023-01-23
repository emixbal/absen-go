package models

import (
	"time"
)

type ClassAttendanceStudent struct {
	MyGorm
	Datang            time.Time       `json:"datang" `
	Pulang            time.Time       `json:"pulang" `
	ClassAttendance   ClassAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassAttendanceID int             `json:"class_attendance_id"`
	Student           Student         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID         int             `json:"student_id"`
}
