package models

import (
	"absen-go/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm"
)

func RecapMemberAttendancePerMember(member_id, year_month string) Response {
	var res Response
	var member Member
	var cams []ClassAttendanceMember
	var member_result MemberResult

	year_month_splited := strings.Split(year_month, "-")
	int_month, _ := strconv.Atoi(year_month_splited[1])
	month_formatted := fmt.Sprintf("%02d", int_month)
	year_formatted := year_month_splited[0]

	start_month, _ := time.Parse("2006-01-02", year_formatted+"-"+month_formatted+"-01")
	end_month := start_month.AddDate(0, 1, -1)

	db := config.GetDBInstance()

	if result := db.Where("is_active = ?", true).First(&member, member_id); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusOK
			res.Message = "Member tidak ditemukan"
			return res
		}

		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong!"
		return res
	}

	db.
		Where("member_id = ?", member.ID).
		Where("arrive > ?", start_month).Where("arrive < ?", end_month).
		Take(&cams)

	var arr_attendances []MemberAttendanceResult
	for _, val := range cams {
		var attendance MemberAttendanceResult

		attendance.Arrive = val.Arrive
		attendance.Leave = val.Leave
		attendance.Date = val.Arrive

		arr_attendances = append(arr_attendances, attendance)
	}

	member_result.ID = member.ID
	member_result.Name = member.Name
	member_result.NIS = member.NIS
	member_result.NISN = member.NISN
	member_result.NBM = member.NBM
	member_result.Attendances = arr_attendances

	res.Data = member_result

	return res
}
