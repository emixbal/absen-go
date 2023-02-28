package models

import "time"

type RoutineOffday struct {
	MyGorm
	Day string `json:"day" gorm:"size:10;not null; index:idx_day"`
}

type Offday struct {
	MyGorm
	Date time.Time `json:"date" gorm:"not null;type:date"`
}
