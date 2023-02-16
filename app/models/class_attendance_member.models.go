package models

import (
	"absen-go/config"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ClassAttendanceMember struct {
	ID uint `json:"id" gorm:"primarykey"`
	// Class    Class     `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	// ClassID  int       `json:"class_id"`
	Member   Member    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	MemberID int       `json:"member_id"`
	Arrive   time.Time `json:"arrive" `
	Leave    time.Time `json:"leave" gorm:"default:null"`
}

func ClassAttendanceMemberArrive(code string) (Response, error) {
	var res Response
	var member Member
	var cam ClassAttendanceMember
	var time_now = time.Now()

	db := config.GetDBInstance()

	// cek apakah member exist
	if result := db.Preload("Class").Where("code = ?", code).First(&member); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "member not found."
			return res, nil
		}
	}

	if result := db.
		Where("DATE(arrive) = ?", time_now.UTC().Format("2006-01-02")).
		Where("member_id = ?", member.ID).
		First(&ClassAttendanceMember{}); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			/**
			Jika belum absen, register simpan absen baru hari ini
			*/
			cam.MemberID = int(member.ID)
			cam.Arrive = time_now
			// cam.ClassID = member.ClassID
			if result := db.Create(&cam); result.Error != nil {
				log.Println("error Create AddClassAttendanceMember")
				log.Println(result.Error)

				res.Status = http.StatusInternalServerError
				res.Message = "error save arrive"
				return res, result.Error
			}

			res.Status = http.StatusOK
			res.Message = "ok"
			res.Data = fiber.Map{
				"member_name":  member.Name,
				"nisn":         member.NISN,
				"class":        member.Class.Name,
				"time_arrival": cam.Arrive.Format("15:04:05"),
			}
			return res, nil
		}
	}

	/**
	Jika sudah absen
	*/
	res.Status = http.StatusBadRequest
	res.Message = "sudah absen"
	return res, nil

}

func ClassAttendanceMemberLeave(code string) (Response, error) {
	var res Response
	var member Member
	// var class_attendance ClassAttendance
	var cam ClassAttendanceMember
	var time_now = time.Now()

	db := config.GetDBInstance()

	if result := db.Preload("Class").Where("code = ?", code).First(&member); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "member not found."
			return res, nil
		}
	}

	if result := db.
		Where("DATE(arrive) = ?", time_now.UTC().Format("2006-01-02")).
		Where("member_id = ?", member.ID).
		First(&cam); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "Belum absen masuk."
			return res, nil
		}
	}

	cam.Leave = time_now
	if result := db.Save(&cam); result.Error != nil {
		log.Println("error Update ClassAttendanceMemberLeave")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update leave"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = fiber.Map{
		"member_name":  member.Name,
		"nisn":         member.NISN,
		"class":        member.Class.Name,
		"time_leaving": cam.Leave.Format("15:04:05"),
	}
	return res, nil
}
