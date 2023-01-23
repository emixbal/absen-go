package models

type Class struct {
	MyGorm
	Name       string   `json:"name" gorm:"size:50"`
	Semester   Semester `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID int      `json:"semester_id"`
}
