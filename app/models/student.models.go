package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name    string `json:"name" gorm:"size:100"`
	NISN    string `json:"nisn" gorm:"size:10"`
	Class   Class  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID int    `json:"class_id"`
}
