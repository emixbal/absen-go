package models

import (
	"time"
)

type ExtracurricularAttendanceStudent struct {
	ID                          uint                      `json:"id" gorm:"primarykey"`
	ExtracurricularAttendance   ExtracurricularAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularAttendanceID int                       `json:"extracurricular_attendance_id"`
	Student                     Student                   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID                   int                       `json:"student_id"`
	Arrive                      time.Time                 `json:"arrive" `
	Leave                       time.Time                 `json:"leave" `
}

func AddExtracurricularAttendanceStudent() (Response, error) {
	var res Response

	return res, nil
}
