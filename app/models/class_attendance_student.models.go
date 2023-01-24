package models

import (
	"absen-go/config"
	"errors"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type ClassAttendanceStudent struct {
	ID                uint            `json:"id" gorm:"primarykey"`
	ClassAttendance   ClassAttendance `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassAttendanceID int             `json:"class_attendance_id"`
	Student           Student         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	StudentID         int             `json:"student_id"`
	Arrive            time.Time       `json:"arrive" `
	Leave             time.Time       `json:"leave" gorm:"default:null"`
}

func AddClassAttendanceStudent(class_attendance_id int, student_id int) (Response, error) {
	var res Response
	var is_not_found bool
	var cas ClassAttendanceStudent

	db := config.GetDBInstance()

	result := db.
		Where("class_attendance_id = ?", class_attendance_id).
		Where("student_id = ?", student_id).
		First(&cas)

	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			is_not_found = true
		}
	}

	if is_not_found {
		cas.ClassAttendanceID = class_attendance_id
		cas.StudentID = student_id
		cas.Arrive = time.Now()

		if result := db.Create(&cas); result.Error != nil {
			log.Println("error Create AddClassAttendanceStudent")
			log.Println(result.Error)

			res.Status = http.StatusBadRequest
			res.Message = "error save arrive"
			return res, result.Error
		}
	} else {
		cas.Leave = time.Now()
		if result := db.Save(&cas); result.Error != nil {
			log.Println("error Update AddClassAttendanceStudent")
			log.Println(result.Error)

			res.Status = http.StatusBadRequest
			res.Message = "error update leave"
			return res, result.Error
		}
	}

	res.Status = http.StatusOK
	res.Message = "success"

	return res, nil

}
