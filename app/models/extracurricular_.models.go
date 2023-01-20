package models

import (
	"time"

	"gorm.io/gorm"
)

type ExtracurricularAttendance struct {
	gorm.Model
	Date                   time.Time            `json:"date" `
	ExtracurricularClass   ExtracurricularClass `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularClassID int                  `json:"extracurricular_class_id"`
}
