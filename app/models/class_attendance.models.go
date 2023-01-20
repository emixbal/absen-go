package models

import (
	"time"

	"gorm.io/gorm"
)

type ClassAttendance struct {
	gorm.Model
	Date    time.Time `json:"date" `
	Class   Class     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID int       `json:"class_id"`
}
