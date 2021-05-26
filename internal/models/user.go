package models

// import "time"
// User struct
type User struct {
	Id               int    `gorm:"primary_key" json:"id"`
	First_name       string `gorm:"not null" json:"first_name"`
	Last_name        string `gorm:"not null" json:"last_name"`
	Phone_number     string `gorm:"not null;unique" json:"phone_number"`
	Username         string `gorm:"not null;unique" json:"username" form:"username"`
	Password         string `gorm:"not null" json:"password" form:"password"`
	Created_at       string `json:"created_at"`
	Updated_at       string `json:"updated_at"`
	Registration_otp string `gorm:"not null" json:"registration_otp"`
}

type UserProfile struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	BankName    string `json:"bank_name"`
	BankAccount string `json:"bank_account"`
	Status      string `json:"status"`
}

// otp formvalue struct
type RegisterFormModel struct {
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Phone_number string `json:"phone_number"`
	Bank_id      int    `json:"bank_id"`
	Bank_account string `json:"bank_account"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Otp          int    `json:"otp"`
}
