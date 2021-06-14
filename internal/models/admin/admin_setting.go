package models

import "time"

// User struct
type Admin_Setting struct {
	Id              int       `gorm:"primary_key" json:"id"`
	DepositWithdraw bool      `gorm:"not null;default:1" json:"deposit_withdraw"`
	Bet             bool      `gorm:"not null;default:1" json:"bet"`
	CancelBet       bool      `gorm:"not null;default:1" json:"cancel_bet"`
	IsActive        bool      `gorm:"not null;default:1" json:"is_active"`
	UsernameBot     string    `json:"username_bot"`
	PasswordBot     string    `json:"password_bot"`
	SelectTextBot   string    `json:"select_text_bot"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`
}
