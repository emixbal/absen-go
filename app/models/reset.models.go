package models

import (
	"absen-go/config"
	"log"
	"net/http"
)

func ResetAttendance() Response {
	var res Response
	var cam ClassAttendanceMember

	db := config.GetDBInstance()
	if result := db.Where("1 = 1").Delete(&cam); result.Error != nil {
		log.Println("ResetAttendance err")
		log.Println(result.Error.Error())
		res.Status = http.StatusInternalServerError
		res.Message = result.Error.Error()
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	return res
}

func ResetMember() Response {
	var res Response

	db := config.GetDBInstance()
	if result := db.Exec("DELETE FROM members"); result.Error != nil {
		log.Println("ResetMember err")
		log.Println(result.Error.Error())
		res.Status = http.StatusInternalServerError
		res.Message = result.Error.Error()
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	return res
}
