package models

type ExtracurricularClass struct {
	MyGorm
	Name              string          `json:"name" gorm:"size:50"`
	Semester          Semester        `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID        int             `json:"semester_id"`
	Student           Student         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID         int             `json:"student_id"`
	Extracurricular   Extracurricular `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularID int             `json:"extracurricular_id"`
}
