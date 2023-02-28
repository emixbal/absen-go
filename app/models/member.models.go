package models

import (
	"absen-go/config"
	"errors"
	"log"
	"math"
	"net/http"

	"gorm.io/gorm"
)

type Member struct {
	MyGorm
	Name     string `json:"name" gorm:"size:100"`
	NISN     string `json:"nisn" gorm:"size:10"`
	NIS      string `json:"nis" gorm:"size:10"`
	NBM      string `json:"nbm" gorm:"size:10"`
	Code     string `json:"code" gorm:"size:20, index:idx_code"`
	Class    Class  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID  int    `json:"class_id"`
	IsActive bool   `json:"is_active,omitempty" gorm:"default:true"`
}

func FethAllMembers(limit int, offset int, class []string, filter_id string) Response {
	type MemberResult struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		NIS       string `json:"nis"`
		NISN      string `json:"nisn"`
		NBM       string `json:"nbm" gorm:"size:10"`
		Code      string `json:"code"`
		ClassName string `json:"class_name"`
		ClassID   string `json:"class_id"`
		IsActive  bool   `json:"is_active"`
	}

	var sudentsResult []MemberResult
	var res Response
	var lr ListResponse
	var total_data int64
	var total_page float64

	db := config.GetDBInstance()

	query := db.Table("members").
		Joins("left join classes on members.class_id = classes.id").
		Select("members.id, members.name, members.nis, members.nisn, members.nbm, members.code, classes.name as class_name, classes.id as class_id, members.is_active").
		Where("members.is_active = ?", true).
		Order("classes.id asc").
		Order("members.name asc")

	if len(class) > 0 {
		query.Where("classes.id IN ?", class)
	}

	if filter_id != "" {
		file_qry := "%" + filter_id + "%"
		query.Where("nis LIKE ?", file_qry).
			Or("nisn LIKE ?", file_qry).
			Or("nbm LIKE ?", file_qry).
			Or("code LIKE ?", file_qry)
	}

	query.Count(&total_data)

	result := query.
		Limit(limit).
		Offset(offset * limit).
		Scan(&sudentsResult)

	total_records := result.RowsAffected

	if result.Error != nil {
		log.Println("Err fetching data")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "Err fetching data"
		return res
	}

	total_page = math.Ceil(float64(total_data) / float64(limit))

	lr.Records = sudentsResult
	lr.TotalData = total_data
	lr.TotalRecords = total_records
	lr.TotalPage = total_page
	lr.Page = offset + 1
	lr.PerPage = limit

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = lr
	return res
}

func MemberUpdate(member_payload *Member, member_id string) Response {
	var res Response
	var member Member

	db := config.GetDBInstance()
	result := db.Where("id = ?", member_id).Take(&member)
	if result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusOK
			res.Message = "can't find member record"
			return res
		}

		log.Panicln("db.Take err")
		log.Panicln(result.Error)
		res.Status = http.StatusInternalServerError
		res.Message = "something went wrong"
		return res
	}

	member.ClassID = member_payload.ClassID
	member.Name = member_payload.Name
	member.NIS = member_payload.NIS
	member.NISN = member_payload.NISN
	member.NBM = member_payload.NBM
	member.Code = member_payload.Code

	if result := db.Save(&member); result.Error != nil {
		log.Panicln("db.Save err")
		log.Panicln(result.Error)
		res.Status = http.StatusInternalServerError
		res.Message = "something went wrong"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = member

	return res
}
