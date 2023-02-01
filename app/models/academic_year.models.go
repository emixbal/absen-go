package models

import (
	"absen-go/config"
	"log"
	"net/http"
)

type AcademicYear struct {
	MyGorm
	Name string `json:"name" gorm:"size:50"`
}

func AcademicYearList() Response {
	var res Response
	var academic_years []AcademicYear

	db := config.GetDBInstance()

	if result := db.Find(&academic_years); result.Error != nil {
		log.Print("error Fetch AcademicYearList")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = academic_years
	return res
}
