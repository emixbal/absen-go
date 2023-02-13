package models

import (
	"absen-go/config"
	"log"
	"net/http"
)

type Class struct {
	MyGorm
	Name     string `json:"name" gorm:"size:50"`
	IsActive bool   `json:"is_active,omitempty" gorm:"default:true"`
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

func ClassAddNew(class *Class) Response {
	var res Response

	db := config.GetDBInstance()
	if result := db.Create(&class); result.Error != nil {
		log.Print("error ClassAddNew")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = class

	return res
}
