package models

import (
	"gorm.io/gorm"
)

type AcademicYear struct {
	gorm.Model
	Name string `json:"name" gorm:"size:50"`
}
