package models

import (
	"time"
)

type ExtracurricularAttendanceMember struct {
	ID                          uint                      `json:"id" gorm:"primarykey"`
	ExtracurricularAttendance   ExtracurricularAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularAttendanceID int                       `json:"extracurricular_attendance_id"`
	Member                      Member                    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	MemberID                    int                       `json:"member_id"`
	Arrive                      time.Time                 `json:"arrive" `
	Leave                       time.Time                 `json:"leave" `
}

func AddExtracurricularAttendanceMember() (Response, error) {
	var res Response

	return res, nil
}
