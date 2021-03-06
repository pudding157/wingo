package models

import "time"

type Password_History struct {
	Id          int       `gorm:"primary_key" json:"id"`
	Username    string    `gorm:"not null" json:"username"`
	IPAddress   string    `json:"ip_address"`
	MACAddress  string    `json:"mac_address"`
	Browser     string    `json:"browser"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	NewPassword string    `json:"new_password"`
	OldPassword string    `json:"old_password"`
}
