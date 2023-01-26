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

func AddClassAttendanceStudent(student_id int) (Response, error) {
	var res Response
	var student Student
	var class_attendance ClassAttendance
	var cas ClassAttendanceStudent
	var is_cas_exist = true // membedakan tap untuk datang atau pulang. Tap 1st untuk datang, selanjutnya dianggap pulang.

	db := config.GetDBInstance()

	if result := db.Preload("Class").Where("id = ?", student_id).First(&student); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "student not found."
			return res, nil
		}
	}

	today := time.Now().UTC().Format("2006-01-02")

	if result := db.
		Where("class_id = ?", student.ClassID).
		Where("date = ?", today).
		First(&class_attendance); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "class not found."
			return res, nil
		}
	}

	if result := db.Where("class_attendance_id = ?", class_attendance.ID).Where("student_id = ?", student.ID).First(&cas); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			is_cas_exist = false
		}
	}

	if !is_cas_exist { // tap pertama (tiba)
		cas.ClassAttendanceID = int(class_attendance.ID)
		cas.StudentID = int(student.ID)
		cas.Arrive = time.Now()

		if result := db.Create(&cas); result.Error != nil {
			log.Println("error Create AddClassAttendanceStudent")
			log.Println(result.Error)

			res.Status = http.StatusInternalServerError
			res.Message = "error save arrive"
			return res, result.Error
		}
	} else { // tap selanjutnya (meninggalkan)
		cas.Leave = time.Now()
		if result := db.Save(&cas); result.Error != nil {
			log.Println("error Update AddClassAttendanceStudent")
			log.Println(result.Error)

			res.Status = http.StatusInternalServerError
			res.Message = "error update leave"
			return res, result.Error
		}
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	return res, nil
}
