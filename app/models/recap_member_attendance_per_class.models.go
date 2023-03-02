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

type MemberAttendanceResult struct {
	Date       time.Time   `json:"date"`
	Arrive     interface{} `json:"arrive"`
	Leave      interface{} `json:"leave"`
	IsAttended bool        `json:"is_attended"`
}

type AttendanceSummary struct {
	TotalDay               uint        `json:"total_day"`
	TotalAbsence           uint        `json:"total_absence"`
	TotalAbsenceWithRemark int         `json:"total_absence_with_remark"`
	TotalSickRemark        interface{} `json:"sick_remarks"`
	TotalOtherkRemark      interface{} `json:"other_remarks"`
}

type MemberResult struct {
	ID                uint                     `json:"id"`
	Name              string                   `json:"name"`
	NIS               string                   `json:"nis"`
	NISN              string                   `json:"nisn"`
	NBM               string                   `json:"nbm"`
	Code              string                   `json:"code"`
	ClassName         string                   `json:"class_name"`
	Attendances       []MemberAttendanceResult `json:"attendances"`
	AttendanceSummary AttendanceSummary        `json:"attendance_summary"`
}

func RecapMemberAttendancePerClass(class_id, year_month string) Response {
	var res Response
	var member_result MemberResult
	var members_result []MemberResult
	var cams []ClassAttendanceMember

	year_month_splited := strings.Split(year_month, "-")
	int_month, _ := strconv.Atoi(year_month_splited[1])
	month_formatted := fmt.Sprintf("%02d", int_month)
	year_formatted := year_month_splited[0]

	start_month, _ := time.Parse("2006-01-02", year_formatted+"-"+month_formatted+"-01")
	end_month := start_month.AddDate(0, 1, -1)

	db := config.GetDBInstance()

	rows_fetch_member, fetch_member_err := db.Table("members").
		Select("members.id, members.name, members.nis, members.nisn, members.nbm, members.code").
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
		var arr_attendances []MemberAttendanceResult
		if err := rows_fetch_member.Scan(&member_result.ID, &member_result.Name, &member_result.NIS, &member_result.NISN, &member_result.NBM, &member_result.Code); err != nil {
			log.Panicln("Err scan members")
			res.Status = http.StatusInternalServerError
			res.Message = "Err scan members"
			return res
		}
		db.
			Where("member_id = ?", member_result.ID).
			Where("arrive > ?", start_month).Where("arrive < ?", end_month).
			Take(&cams)

		for _, val := range cams {
			var attendance MemberAttendanceResult

			attendance.Date = val.Arrive

			attendance.Arrive = val.Arrive

			if val.Leave.Before(val.Arrive) {
				attendance.Leave = nil
			} else {
				attendance.Leave = val.Leave
			}

			arr_attendances = append(arr_attendances, attendance)
		}

		member_result.Attendances = arr_attendances
		members_result = append(members_result, member_result)
	}

	if len(members_result) < 1 {
		res.Status = http.StatusOK
		res.Message = "no data"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = members_result
	return res
}
