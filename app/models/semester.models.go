package models

import (
	"gorm.io/gorm"
)

type Semester struct {
	gorm.Model
	Name           string       `json:"name" gorm:"size:50"`
	AcademicYear   AcademicYear `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	AcademicYearID int          `json:"academic_year_id"`
}
