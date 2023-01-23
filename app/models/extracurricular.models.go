package models

type Extracurricular struct {
	MyGorm
	Name       string   `json:"name" gorm:"size:50"`
	Teacher    Teacher  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	TeacherID  int      `json:"teacher_id"`
	Semester   Semester `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID int      `json:"semester_id"`
}
