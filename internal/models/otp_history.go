package models

import "time"

// Otp struct
type Otp_History struct {
	Id        int       `gorm:"primary_key" json:"id"`
	Type      int       `json:"type"`
	SendTo    string    `json:"send_to"`
	Otp       int       `json:"otp"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
}
