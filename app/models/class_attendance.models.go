package models

import (
	"time"
)

type ClassAttendance struct {
	MyGorm
	Date    time.Time `json:"date" `
	Class   Class     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID int       `json:"class_id"`
}
