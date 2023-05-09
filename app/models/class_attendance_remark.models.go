package models

import (
	"absen-go/config"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type ClassAttendanceRemarkType struct {
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"size:50"`
}

type ClassAttendanceRemark struct {
	ID           uint                      `json:"id" gorm:"primarykey"`
	Member       Member                    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	MemberID     int                       `json:"member_id"`
	Date         time.Time                 `json:"date"`
	RemarkType   ClassAttendanceRemarkType `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	RemarkTypeID int                       `json:"remark_type_id"`
	Text         string                    `json:"text" gorm:"type:text"`
}

func ClassAttendanceRemarking(code string, remark_type_id int, text string, t time.Time) Response {
	var res Response
	var member Member
	var car ClassAttendanceRemark
	var cam ClassAttendanceMember

	db := config.GetDBInstance()

	// cek apakah member exist
	if result := db.Preload("Class").Where("code = ?", code).First(&member); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "member not found."
			return res
		}
	}

	//cek apakah sudah absen
	if result := db.Where("DATE(arrive) = ?", t.UTC().Format("2006-01-02")).Take(&cam); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			fmt.Println("===== ok, emang belum absen =====")
		}
	} else {
		res.Status = http.StatusBadRequest
		res.Message = "Sudah absen masuk"
		return res
	}

	//cek apakah sudah absen
	if result := db.
		Where("DATE(date) = ?", t.UTC().Format("2006-01-02")).
		Where("member_id = ?", member.ID).
		Take(&car); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			fmt.Println("===== ok, emang belum dibuat keterangan =====")
		}
	} else {
		res.Status = http.StatusBadRequest
		res.Message = "Keterangan untuk user dan hari ini telah dibuat."
		return res
	}

	car.MemberID = int(member.ID)
	car.Date = t
	car.RemarkTypeID = remark_type_id
	car.Text = text

	if result := db.Create(&car); result.Error != nil {
		log.Print("error ClassAttendanceRemarking")
		log.Print(result.Error)

		res.Status = http.StatusBadRequest
		res.Message = "error"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = car

	return res
}
