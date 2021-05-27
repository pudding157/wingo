package models

// import "time"
// Otp struct
type Otp_History struct {
	Id         int    `gorm:"primary_key" json:"id"`
	Type       int    `json:"type"`
	Send_to    string `json:"send_to"`
	Otp        int    `json:"otp"`
	Created_at string `json:"created_at"`
}
