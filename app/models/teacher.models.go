package models

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name string `json:"name" gorm:"size:50"`
	NIP  string `json:"nip" gorm:"size:18"`
}
