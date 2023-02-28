package models

import "time"

type RoutineOffday struct {
	MyGorm
	Day string `json:"day" gorm:"size:10;not null"`
}

type Offday struct {
	MyGorm
	Date time.Time `json:"date" gorm:"not null;type:date"`
}
