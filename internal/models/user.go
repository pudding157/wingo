package models

import "time"

// User struct
type User struct {
	Id              int            `gorm:"primary_key" json:"id"`
	FirstName       string         `gorm:"not null" json:"first_name"`
	LastName        string         `gorm:"not null" json:"last_name"`
	PhoneNumber     string         `gorm:"not null;unique" json:"phone_number"`
	Username        string         `gorm:"not null;unique" json:"username" form:"username"`
	Password        string         `gorm:"not null" json:"password" form:"password"`
	CreatedAt       time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"not null" json:"updated_at"`
	RegistrationOtp string         `gorm:"not null" json:"registration_otp"`
	Status          int            `gorm:"not null;default:0" json:"status"`
	Affiliate       string         `gorm:"" json:"affiliate"`
	ParentUserId    *int           `gorm:"" json:"parent_user_id"`
	UserHistory     []User_History `gorm:"ForeignKey:UserId"`
	// User_Wallet     User_Wallet `json:"User_Wallet"`
}

type UserProfile struct {
	Username       string   `json:"username"`
	Name           string   `json:"full_name"`
	PhoneNumber    string   `json:"phone_number"`
	BankName       string   `json:"bank_name"`
	BankAccount    string   `json:"bank_account"`
	Status         string   `json:"status"`
	ParentUserName *string  `json:"parent_username"`
	ChildUserNames []string `json:"child_usernames"`
}

type RegisterFormModel struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	BankId        int    `json:"bank_id"`
	BankAccount   string `json:"bank_account"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Otp           int    `json:"otp"`
	AffiliateCode string `json:"affiliate_code"`
}
