package models

// import "time"
// User struct
type User struct {
	Id               int    `gorm:"primary_key" json:"id"`
	First_name       string `json:"first_name"`
	Last_name        string `json:"last_name"`
	Phone_number     string `json:"phone_number"`
	Bank_id          int    `json:"bank_id"`
	Bank_account     string `json:"bank_account"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Created_at       string `json:"created_at"`
	Updated_at       string `json:"updated_at"`
	Registration_otp string `json:"registration_otp"`
}
