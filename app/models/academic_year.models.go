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
	Name string `json:"name" gorm:"size:50, index:idx_name,unique"`
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
	res.Message = "ok"
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
	res.Message = "ok"
	res.Data = academic_year
	return res
}

func AcademicYearUpdate(year_id string, new_name string) Response {
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

	academic_year.Name = new_name

	if result := db.Save(&academic_year); result.Error != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong!"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = academic_year
	return res
}

func AcademicYearNew(new_name string) Response {
	var res Response
	var academic_year AcademicYear
	var isNameExist = true

	academic_year.Name = new_name

	db := config.GetDBInstance()
	if result := db.Where("name = ?", new_name).Take(&academic_year); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			isNameExist = false
		}

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res
	}

	if isNameExist {
		res.Status = http.StatusBadRequest
		res.Message = "name already exist"
		return res
	}

	if result := db.Create(&academic_year); result.Error != nil {
		log.Print("error AcademicYearNew")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = academic_year
	return res
}
