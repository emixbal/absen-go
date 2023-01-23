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
}
