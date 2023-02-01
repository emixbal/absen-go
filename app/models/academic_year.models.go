package models

import (
	"absen-go/config"
	"errors"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
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

func AcademicYearDetail(year_id string) Response {
	var res Response
	var academic_year AcademicYear

	_, err := strconv.Atoi(year_id)
	if err != nil {
		log.Println(err)
		res.Status = http.StatusBadRequest
		res.Message = "invalid id"
		return res
	}

	db := config.GetDBInstance()
	result := db.Where("deleted_at IS NULL").First(&academic_year, year_id)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusOK
			res.Message = "can't find record"
			return res
		}

		log.Print("error Get AcademicYearDetail")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong!"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = academic_year
	return res
}
