package models

import (
	"absen-go/config"
	"log"
	"net/http"
)

type Class struct {
	MyGorm
	Name       string   `json:"name" gorm:"size:50"`
	IsActive   bool     `json:"is_active,omitempty" gorm:"default:true"`
	Semester   Semester `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	SemesterID int      `json:"semester_id"`
}

func ClassFetchAll() Response {
	var classes []Class
	var res Response

	db := config.GetDBInstance()

	if result := db.Where("is_active = ?", true).Find(&classes); result.Error != nil {
		log.Print("error ClassFetchAll")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = classes

	return res
}
