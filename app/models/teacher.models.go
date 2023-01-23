package models

type Teacher struct {
	MyGorm
	Name string `json:"name" gorm:"size:50"`
	NIP  string `json:"nip" gorm:"size:18"`
}
