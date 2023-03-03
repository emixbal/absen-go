package models

import (
	"absen-go/config"
	"log"
	"net/http"
	"time"
)

type RoutineOffday struct {
	MyGorm
	Day string `json:"day" gorm:"size:10;not null; index:idx_day"`
}

type Offday struct {
	MyGorm
	Date time.Time `json:"date" gorm:"not null;type:date; index:idx_date"`
}

func OffdayAddNew(offday *Offday) Response {
	var res Response
	var count int64

	db := config.GetDBInstance()

	db.Find(&offday).Count(&count)
	if count > 1 {
		res.Status = http.StatusBadRequest
		res.Message = "Tanngal tersebut sudah ditambahkan sebagai hari libur "
		return res
	}

	if result := db.Create(&offday); result.Error != nil {
		log.Print("error OffdayAddNew")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error OffdayAddNew"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = offday

	return res
}

func OffdayFetchAll() (Response, error) {
	var offdays []Offday
	var res Response

	db := config.GetDBInstance()

	if result := db.Find(&offdays); result.Error != nil {
		log.Print("error OffdayFetchAll")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = offdays

	return res, nil
}
