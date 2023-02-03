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

func FethAllStudents(limit int, offset int, class []string) Response {
	type StudentResult struct {
		ID        string `json:"id"`
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
		Joins("left join classes on students.class_id = classes.id").
		Select("students.id, students.name, students.nisn, classes.name as class_name, classes.id as class_id, students.is_active").
		Where("students.is_active = ?", true).
		Order("classes.id asc").
		Order("students.name asc")

	if len(class) > 0 {
		query.Where("classes.id IN ?", class)
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
