package models

import (
	"time"

	"gorm.io/gorm"
)

type Remark struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	Member      Member         `gorm:"constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT;-"`
	MemberId    uint           `json:"member_id"`
	Date        time.Time      `json:"date" gorm:"not null;type:date"`
	IsSick      bool           `json:"is_sick"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
