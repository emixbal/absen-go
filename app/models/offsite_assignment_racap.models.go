package models

import (
	"absen-go/config"
	"log"
	"net/http"
	"time"
)

func OffSiteAssignmentRecapAll(sort_by, sort_type string, start_date, end_date time.Time) Response {
	type Result struct {
		ID          uint      `json:"id"`
		Name        string    `json:"name"`
		Class       string    `json:"class"`
		NIS         string    `json:"nis"`
		NISN        string    `json:"nisn"`
		NBM         string    `json:"nbm"`
		Description string    `json:"description"`
		Departure   time.Time `json:"departure" `
		Arrive      time.Time `json:"arrive"`
	}

	var res Response
	var results []Result

	db := config.GetDBInstance()

	q := db.Table("off_site_assignment_members osam").
		Select("osam.id, m.name, c.name as class, m.nis, m.nisn, m.nbm, osam.description, osam.departure, osam.arrive").
		Joins("left join members m on osam.member_id = m.id").
		Joins("left join classes c on m.class_id = c.id").
		Where("osam.departure > ?", start_date).Where("osam.departure < ?", end_date)

	if sort_by == "class" {
		q.Order("c.id " + sort_type)
	}

	if sort_by == "departure" {
		q.Order("osam.departure " + sort_type)
	}

	if sort_by == "name" {
		q.Order("m.name " + sort_type)
	}

	result := q.Scan(&results)

	if result.Error != nil {
		log.Print("error OffSiteAssignmentRecapAll")
		log.Print(result.Error)

		res.Status = http.StatusInternalServerError
		res.Message = "error fetchin records"
		return res
	}

	res.Status = http.StatusOK
	res.Message = "ok"
	res.Data = results

	return res
}
