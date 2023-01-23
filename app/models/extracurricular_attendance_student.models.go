package models

import (
	"time"

	"gorm.io/gorm"
)

type ExtracurricularAttendanceStudent struct {
	gorm.Model
	Datang                      time.Time                 `json:"datang" `
	Pulang                      time.Time                 `json:"pulang" `
	ExtracurricularAttendance   ExtracurricularAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularAttendanceID int                       `json:"extracurricular_attendance_id"`
	Student                     Student                   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID                   int                       `json:"student_id"`
}
