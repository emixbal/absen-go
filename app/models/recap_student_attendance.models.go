package models

import (
	"absen-go/config"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RecapStudentAttendance(class_id, month string) Response {
	type StudentResult struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		NISN string `json:"nisn"`
	}

	var res Response
	var class_attendances []ClassAttendance
	var sudents_result []StudentResult
	db := config.GetDBInstance()

	if result := db.Table("students").
		Select("students.id, students.name, students.nisn, students.code").
		Where("students.class_id = ?", class_id).
		Order("students.name asc").
		Scan(&sudents_result); result.Error != nil {

		log.Println("Err fetching data")
		log.Println(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "Err fetching students data"
		return res
	}

	start_month, _ := time.Parse("2006-01-02", "2023-"+month+"-01")
	end_month, _ := time.Parse("2006-01-02", "2023-"+month+"-28")

	if result := db.Where("date > ?", start_month).Where("date < ?", end_month).Find(&class_attendances); result.Error != nil {
		log.Print("error fetch class_attendances")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin class_attendances data"
		return res
	}

	if len(class_attendances) < 1 {
		res.Status = http.StatusOK
		res.Message = "tidak ada hari kerja pada bulan " + month
	}

	fmt.Println(class_attendances)

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = fiber.Map{
		"student_in_class":  sudents_result,
		"class_attendances": class_attendances,
	}

	return res
}
