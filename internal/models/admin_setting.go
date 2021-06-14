package models

import "time"

// User struct
type Admin_Setting struct {
	Id              int       `gorm:"primary_key" json:"id"`
	FirstName       string    `gorm:"not null" json:"first_name"`
	LastName        string    `gorm:"not null" json:"last_name"`
	PhoneNumber     string    `gorm:"not null;unique" json:"phone_number"`
	Username        string    `gorm:"not null;unique" json:"username" form:"username"`
	Password        string    `gorm:"not null" json:"password" form:"password"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`
	RegistrationOtp string    `gorm:"not null" json:"registration_otp"`
	Status          int       `gorm:"not null;default:0" json:"status"`
	Affiliate       string    `gorm:"" json:"affiliate"`
	ParentUserId    *int      `gorm:"" json:"parent_user_id"`
}
