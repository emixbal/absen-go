package models

import (
	"absen-go/config"
	"log"
	"net/http"
	"time"
)

type ExtracurricularAttendanceStudent struct {
	ID                          uint                      `json:"id" gorm:"primarykey"`
	ExtracurricularAttendance   ExtracurricularAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ExtracurricularAttendanceID int                       `json:"extracurricular_attendance_id"`
	Student                     Student                   `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID                   int                       `json:"student_id"`
	Arrive                      time.Time                 `json:"arrive" `
	Leave                       time.Time                 `json:"leave" `
}

func AddExtracurricularAttendanceStudent(book *Book) (Response, error) {
	var res Response
	db := config.GetDBInstance()

	if result := db.Create(&book); result.Error != nil {
		log.Println("error AddExtracurricularAttendanceStudent")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error save new record"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = book

	return res, nil
}
