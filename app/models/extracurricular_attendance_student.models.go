package models

import "time"

type ExtracurricularAttendanceStudent struct {
	ID                          uint                      `json:"id" gorm:"primarykey"`
	ExtracurricularAttendance   ExtracurricularAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularAttendanceID int                       `json:"extracurricular_attendance_id"`
	Student                     Student                   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID                   int                       `json:"student_id"`
	Datang                      time.Time                 `json:"datang" `
	Pulang                      time.Time                 `json:"pulang" `
}
