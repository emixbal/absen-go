package models

import (
	"time"
)

type ClassAttendance struct {
	MyGorm
	Date       time.Time `json:"date" gorm:"type:date(3)"`
	Class      Class     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID    uint      `json:"class_id"`
	Semester   Semester  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID uint      `json:"semester_id"`
	StartTime  string    `gorm:"type:time(3)" json:"start_time,omitempty"`
	EndTime    string    `gorm:"type:time(3)" json:"end_time,omitempty"`
}
