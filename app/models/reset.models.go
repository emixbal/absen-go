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
		log.Println("Delete err")
		log.Println(result.Error.Error())
		res.Status = http.StatusInternalServerError
		res.Message = result.Error.Error()
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	return res
}
