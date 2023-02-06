package models

import (
	"absen-go/config"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func RecapStudentAttendance(class_id, month string) Response {
	type StudentAttendance struct {
		AttendanceID int       `json:"attendance_id"`
		Arrive       time.Time `json:"arrive"`
		Leave        time.Time `json:"leave"`
	}
	type StudentResult struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		NISN        string `json:"nisn"`
		Code        string `json:"code"`
		Attendances []StudentAttendance
	}

	var res Response
	var sudent_result StudentResult
	var sudents_result []StudentResult
	var class_attendances []ClassAttendance

	db := config.GetDBInstance()

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
		return res
	}

	rows_fetch_student, fetch_student_err := db.Table("students").
		Select("students.id, students.name, students.nisn, students.code").
		Where("students.class_id = ?", class_id).
		Order("students.name asc").Rows()

	defer rows_fetch_student.Close()

	if fetch_student_err != nil {
		log.Println("Err fetching data")
		log.Println(fetch_student_err)

		res.Status = http.StatusInternalServerError
		res.Message = "Err fetching students data"
		return res
	}

	for rows_fetch_student.Next() {
		if err := rows_fetch_student.Scan(&sudent_result.ID, &sudent_result.Name, &sudent_result.NISN, &sudent_result.Code); err != nil {
			res.Status = http.StatusInternalServerError
			res.Message = "Err scan students"
			return res
		}

		var cases []ClassAttendanceStudent
		for _, class_attendance := range class_attendances {
			var cas ClassAttendanceStudent
			db.Where("class_attendance_id = ?", class_attendance.ID).Where("student_id = ?", sudent_result.ID).Take(&cas)
			cases = append(cases, cas)
		}

		var arr_cas []StudentAttendance
		for _, val := range cases {
			var cas StudentAttendance

			cas.AttendanceID = val.ClassAttendanceID
			cas.Arrive = val.Arrive
			cas.Leave = val.Leave

			arr_cas = append(arr_cas, cas)
		}

		sudent_result.Attendances = arr_cas

		sudents_result = append(sudents_result, sudent_result)
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = fiber.Map{
		"student_in_class": sudents_result,
	}
	return res
}
