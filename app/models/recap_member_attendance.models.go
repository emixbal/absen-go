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

func RecapMemberAttendance(class_id, year_month string) Response {
	type MemberAttendance struct {
		AttendanceID int       `json:"attendance_id"`
		Date         time.Time `json:"date"`
		Arrive       time.Time `json:"arrive"`
		Leave        time.Time `json:"leave"`
	}
	type MemberResult struct {
		ID          string             `json:"id"`
		Name        string             `json:"name"`
		NIS         string             `json:"nis"`
		NISN        string             `json:"nisn"`
		Code        string             `json:"code"`
		Attendances []MemberAttendance `json:"attendances"`
	}

	var res Response
	var sudent_result MemberResult
	var sudents_result []MemberResult
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

	rows_fetch_member, fetch_member_err := db.Table("members").
		Select("members.id, members.name, members.nis, members.nisn, members.code").
		Where("members.class_id = ?", class_id).
		Order("members.name asc").Rows()

	defer rows_fetch_member.Close()

	if fetch_member_err != nil {
		log.Println("Err fetching data")
		log.Println(fetch_member_err)

		res.Status = http.StatusInternalServerError
		res.Message = "Err fetching members data"
		return res
	}

	for rows_fetch_member.Next() {
		if err := rows_fetch_member.Scan(&sudent_result.ID, &sudent_result.Name, &sudent_result.NIS, &sudent_result.NISN, &sudent_result.Code); err != nil {
			log.Panicln("Err scan members")
			res.Status = http.StatusInternalServerError
			res.Message = "Err scan members"
			return res
		}

		var cases []ClassAttendanceMember
		for _, class_attendance := range class_attendances {
			var cas ClassAttendanceMember
			db.Preload("ClassAttendance").Where("class_attendance_id = ?", class_attendance.ID).Where("member_id = ?", sudent_result.ID).Take(&cas)

			cases = append(cases, cas)
		}

		var arr_cas []MemberAttendance
		for _, val := range cases {
			var cas MemberAttendance

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