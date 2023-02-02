package models

import (
	"absen-go/config"
	"log"
	"math"
	"net/http"
)

type Student struct {
	MyGorm
	Name     string `json:"name" gorm:"size:100"`
	NISN     string `json:"nisn" gorm:"size:10"`
	Class    Class  `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;"`
	ClassID  int    `json:"class_id"`
	IsActive bool   `json:"is_active,omitempty" gorm:"default:true"`
}

func FethAllStudents(limit int, offset int, class string) Response {
	type StudentResult struct {
		Name      string `json:"name"`
		NISN      string `json:"nisn"`
		ClassName string `json:"class_name"`
		ClassID   string `json:"class_id"`
		IsActive  bool   `json:"is_active"`
	}

	var sudentsResult []StudentResult
	var res Response
	var lr ListResponse
	var total_data int64
	var total_page float64

	db := config.GetDBInstance()

	query := db.Table("students").
		Select("students.name, students.nisn, classes.name as class_name, classes.id as class_id, students.is_active").
		Joins("left join classes on students.class_id = classes.id").
		Where("students.is_active = ?", true).
		Order("classes.name asc").
		Order("students.name asc")

	if class != "" {
		query.Where("classes.name = ?", class)
	}

	query.Count(&total_data)

	result := query.
		Limit(limit).
		Offset(offset).
		Scan(&sudentsResult)

	if result.Error != nil {
		log.Println("Err fetching data")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "Err fetching data"
		return res
	}

	total_page = math.Ceil(float64(total_data) / float64(limit))

	lr.Record = sudentsResult
	lr.TotalData = total_data
	lr.TotalPage = total_page
	lr.Page = offset + 1
	lr.PerPage = limit

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = lr
	return res
}
