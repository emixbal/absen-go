package models

import (
	"time"
)

type ClassAttendance struct {
	MyGorm
	Date       time.Time `json:"date" `
	Class      Class     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID    uint      `json:"class_id"`
	Semester   Semester  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID uint      `json:"semester_id"`
	StartTime  time.Time `gorm:"type:time(3)" json:"start_time,omitempty"`
	EndTime    time.Time `gorm:"type:time(3)" json:"end_time,omitempty"`
}
