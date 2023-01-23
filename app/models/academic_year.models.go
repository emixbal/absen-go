package models

type AcademicYear struct {
	MyGorm
	Name string `json:"name" gorm:"size:50"`
}
