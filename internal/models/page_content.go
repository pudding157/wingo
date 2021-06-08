package models

import "time"

// Page Content struct
type Page_Content struct {
	Id          int       `gorm:"primary_key" json:"id"`
	RunningText string    `json:"running_text"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}
