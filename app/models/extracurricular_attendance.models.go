package models

import (
	"time"
)

type ExtracurricularAttendance struct {
	MyGorm
	Date                   time.Time            `json:"date" `
	ExtracurricularClass   ExtracurricularClass `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularClassID int                  `json:"extracurricular_class_id"`
}
