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

type Remark struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Member      Member         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;-"`
	MemberId    uint           `json:"member_id"`
	Date        time.Time      `json:"date" gorm:"not null;type:date"`
	IsSick      bool           `json:"is_sick"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func ClassAttendanceRemarking(code string, is_sick bool, text string, t time.Time) Response {
	var res Response
	var member Member
	var remark Remark
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
		Take(&remark); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			fmt.Println("===== ok, emang belum dibuat keterangan =====")
		}
	} else {
		res.Status = http.StatusBadRequest
		res.Message = "Keterangan untuk user dan hari ini telah dibuat."
		return res
	}

	remark.MemberId = uint(member.ID)
	remark.Date = t
	remark.Description = text
	remark.IsSick = is_sick

	if result := db.Create(&remark); result.Error != nil {
		log.Print("error ClassAttendanceRemarking")
		log.Print(result.Error)

		res.Status = http.StatusBadRequest
		res.Message = "error"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = remark

	return res
}
