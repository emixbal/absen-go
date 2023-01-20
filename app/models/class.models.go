package models

import (
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name       string   `json:"name" gorm:"size:50"`
	Semester   Semester `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID int      `json:"semester_id"`
}
