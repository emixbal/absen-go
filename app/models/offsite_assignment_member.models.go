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

type OffSiteAssignmentMember struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Member      Member    `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	MemberID    int       `json:"member_id"`
	Description string    `json:"description" gorm:"type:text"`
	Departure   time.Time `json:"departure" `
	Arrive      time.Time `json:"arrive" gorm:"default:null"`
}

func OffSiteAssignmentDeparture(code, description string) (Response, error) {
	var res Response
	var member Member
	var oad OffSiteAssignmentMember
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
		First(&cam); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			/**
			Jika belum absen, proses return. Yang tugas keluar harus sudah absen terlebih dahulu.
			*/
			res.Status = http.StatusBadRequest
			res.Message = "Belum absen kedatangan"
			return res, nil
		}
	}

	if result := db.
		Where("DATE(departure) = ?", time_now.UTC().Format("2006-01-02")).
		Where("arrive IS NULL").
		Where("member_id = ?", member.ID).
		First(&oad); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			/**
			Jika belum absen, register simpan absen baru hari ini
			*/
			oad.MemberID = int(member.ID)
			oad.Departure = time_now
			oad.Description = description
			if result := db.Create(&oad); result.Error != nil {
				log.Println("error Create OffSiteAssignmentDeparture")
				log.Println(result.Error)

				res.Status = http.StatusInternalServerError
				res.Message = "error save departure"
				return res, result.Error
			}

			res.Status = http.StatusOK
			res.Message = "ok"
			res.Data = fiber.Map{
				"member_name":    member.Name,
				"nisn":           member.NISN,
				"nbm":            member.NBM,
				"class":          member.Class.Name,
				"time_departure": oad.Departure.Format("15:04:05"),
				"description":    oad.Description,
			}
			return res, nil
		}
	}

	/**
	Jika sudah ada penugasan dihari yang sama dan belum melakukan absen kembali. Tidak bisa dua tugas di waktu yang sama.
	*/
	res.Status = http.StatusBadRequest
	res.Message = "sudah tugas keluar"
	return res, nil
}

func OffSiteAssignmentArrive(code string) (Response, error) {
	var res Response
	var member Member
	var oad OffSiteAssignmentMember
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
		Where("DATE(departure) = ?", time_now.UTC().Format("2006-01-02")).
		Where("arrive IS NULL").
		Where("member_id = ?", member.ID).
		First(&oad); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusBadRequest
			res.Message = "Tidak dalam tugas/ijin."
			return res, nil
		}
	}

	oad.Arrive = time_now
	if result := db.Save(&oad); result.Error != nil {
		log.Println("error Update OffSiteAssignmentArrive")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error update arrive"
		return res, result.Error
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = fiber.Map{
		"member_name":   member.Name,
		"nisn":          member.NISN,
		"nbm":           member.NBM,
		"class":         member.Class.Name,
		"time_arriving": oad.Arrive.Format("15:04:05"),
	}

	return res, nil
}
