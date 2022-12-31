package models

type Victim struct {
	UniqueID string `json:"unique_id" gorm:"primary_key"`
	Key      string `json:"key"`
}
