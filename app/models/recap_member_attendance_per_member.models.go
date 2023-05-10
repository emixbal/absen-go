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
	type RemarkResult struct {
		Date        string `json:"date"`
		Description string `json:"description"`
	}

	var res Response
	var member Member
	var cam ClassAttendanceMember
	var member_result MemberResult
	var attendance MemberAttendanceResult
	var attendance_summary AttendanceSummary

	year_month_splited := strings.Split(year_month, "-")
	int_month, _ := strconv.Atoi(year_month_splited[1])
	month_formatted := fmt.Sprintf("%02d", int_month)
	year_formatted := year_month_splited[0]

	start_month, _ := time.Parse("2006-01-02", year_formatted+"-"+month_formatted+"-01")
	end_month := start_month.AddDate(0, 1, -1)

	db := config.GetDBInstance()

	if result := db.Preload("Class").Where("is_active = ?", true).First(&member, member_id); result.Error != nil {
		if is_notfound := errors.Is(result.Error, gorm.ErrRecordNotFound); is_notfound {
			res.Status = http.StatusOK
			res.Message = "Member tidak ditemukan"
			return res
		}

		res.Status = http.StatusInternalServerError
		res.Message = "Something went wrong!"
		return res
	}

	/**
	FETCHING HARI LIBUR MINGGUAN START
	*/
	routine_offdays := []string{}
	db.
		Table("routine_offdays").
		Select("routine_offdays.day").
		Scan(&routine_offdays)
	/**
	FETCHING HARI LIBUR MINGGUAN END
	*/
	/**


	FETCHING HARI LIBUR INSIDENTIL START
	*/
	offdays := []time.Time{}
	db.
		Table("offdays").
		Select("offdays.date").
		Scan(&offdays)
	/**
	FETCHING HARI LIBUR INSIDENTIL END
	*/

	var arr_attendances []MemberAttendanceResult
	var total_absence = uint(0)
	var total_day = uint(0)

	for d := start_month; !d.After(end_month); d = d.AddDate(0, 0, 1) {
		if sliceOfStringContains(routine_offdays, d.Weekday().String()) { // check apakah hari d adalah hari libur mingguan?
			attendance.IsOffday = true
			fmt.Println("=====bukan hari masuk=====")
		} else {
			if sliceOfTimeContainDate(offdays, d) {
				attendance.IsOffday = true
				fmt.Println("=====hari libur=====")
			} else {
				cam = ClassAttendanceMember{}
				result := db.
					Where("member_id = ?", member.ID).
					Where("arrive > ? AND arrive < ?", d, d.Add(24*time.Hour)).
					First(&cam)

				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
					attendance.IsAttended = false
					attendance.IsOffday = false
					attendance.Arrive = nil
					attendance.Leave = nil
					total_absence++
				} else {
					attendance.IsAttended = true
					attendance.IsOffday = false
					attendance.Arrive = cam.Arrive
					if cam.Leave.Before(cam.Arrive) {
						attendance.Leave = nil
					} else {
						attendance.Leave = cam.Leave
					}
				}
				total_day++
			}
		}
		attendance.Date = d
		arr_attendances = append(arr_attendances, attendance)
	}

	var sick_remark_result []RemarkResult
	var other_remark_result []RemarkResult
	db.
		Table("remarks").
		Select("remarks.date, remarks.description").
		Where("remarks.is_sick = ?", true).
		Where("remarks.date > ? AND remarks.date < ?", start_month.Format("2006-01-02"), end_month.Format("2006-01-02")).
		Scan(&sick_remark_result)

	db.
		Table("remarks").
		Select("remarks.date, remarks.description").
		Where("remarks.is_sick = ?", false).
		Where("remarks.date > ? AND remarks.date < ?", start_month.Format("2006-01-02"), end_month.Format("2006-01-02")).
		Scan(&other_remark_result)

	attendance_summary.TotalAbsence = total_absence
	attendance_summary.TotalDay = total_day
	attendance_summary.SickRemark = sick_remark_result
	attendance_summary.TotalSick = len(sick_remark_result)
	attendance_summary.TotalOtherkRemark = len(other_remark_result)
	attendance_summary.OtherkRemark = other_remark_result
	attendance_summary.TotalAbsenceWithRemark = len(sick_remark_result) + len(other_remark_result)
	attendance_summary.TotalAbsenceNoRemark = int(total_absence) - attendance_summary.TotalAbsenceWithRemark

	member_result.ID = member.ID
	member_result.Name = member.Name
	member_result.NIS = member.NIS
	member_result.NISN = member.NISN
	member_result.NBM = member.NBM
	member_result.Code = member.Code
	member_result.ClassName = member.Class.Name
	member_result.Attendances = arr_attendances
	member_result.AttendanceSummary = attendance_summary

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = member_result

	return res
}

func sliceOfStringContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func sliceOfTimeContainDate(dates []time.Time, date time.Time) bool {
	for _, d := range dates {
		if d.Truncate(24 * time.Hour).Equal(date.Truncate(24 * time.Hour)) {
			return true
		}
	}

	return false
}
