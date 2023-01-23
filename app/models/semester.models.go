package models

type Semester struct {
	MyGorm
	Name           string       `json:"name" gorm:"size:1"`
	AcademicYear   AcademicYear `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	AcademicYearID int          `json:"academic_year_id"`
}
