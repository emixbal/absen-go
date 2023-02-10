package models

import (
	"absen-go/config"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RecapStudentAttendance(class_id, year_month string) Response {
	type StudentAttendance struct {
		AttendanceID int       `json:"attendance_id"`
		Date         time.Time `json:"date"`
		Arrive       time.Time `json:"arrive"`
		Leave        time.Time `json:"leave"`
	}
	type StudentResult struct {
		ID          string              `json:"id"`
		Name        string              `json:"name"`
		NISN        string              `json:"nisn"`
		Code        string              `json:"code"`
		Attendances []StudentAttendance `json:"attendances"`
	}

	var res Response
	var sudent_result StudentResult
	var sudents_result []StudentResult
	var class_attendances []ClassAttendance

	year_month_splited := strings.Split(year_month, "-")

	int_month, _ := strconv.Atoi(year_month_splited[1])
	month_formatted := fmt.Sprintf("%02d", int_month)
	year_formatted := year_month_splited[0]

	start_month, _ := time.Parse("2006-01-02", year_formatted+"-"+month_formatted+"-01")
	end_month := start_month.AddDate(0, 1, -1)

	db := config.GetDBInstance()

	if result := db.Where("date > ?", start_month).Where("date < ?", end_month).Find(&class_attendances); result.Error != nil {
		log.Print("error fetch class_attendances")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin class_attendances data"
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
			db.Preload("ClassAttendance").Where("class_attendance_id = ?", class_attendance.ID).Where("student_id = ?", sudent_result.ID).Take(&cas)

			cases = append(cases, cas)
		}

		var arr_cas []StudentAttendance
		for _, val := range cases {
			var cas StudentAttendance

			cas.AttendanceID = val.ClassAttendanceID
			cas.Arrive = val.Arrive
			cas.Leave = val.Leave
			cas.Date = val.ClassAttendance.Date

			arr_cas = append(arr_cas, cas)
		}

		sudent_result.Attendances = arr_cas

		sudents_result = append(sudents_result, sudent_result)
	}

	if len(sudents_result) < 1 {
		res.Status = http.StatusOK
		res.Message = "no data"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = sudents_result
	return res
}
