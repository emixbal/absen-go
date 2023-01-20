package models

import (
	"gorm.io/gorm"
)

type Extracurricular struct {
	gorm.Model
	Name       string   `json:"name" gorm:"size:50"`
	Teacher    Teacher  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	TeacherID  int      `json:"teacher_id"`
	Semester   Semester `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID int      `json:"semester_id"`
}
