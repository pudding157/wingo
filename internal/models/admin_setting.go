package models

import "time"

// Admin_Setting struct
type Admin_Setting struct {
	Id              int       `gorm:"primary_key" json:"id"`
	DepositWithdraw bool      `gorm:"not null;default:1" json:"deposit_withdraw"`
	Bet             bool      `gorm:"not null;default:1" json:"bet"`
	CancelBet       bool      `gorm:"not null;default:1" json:"cancel_bet"`
	IsActive        bool      `gorm:"not null;default:1" json:"is_active"`
	UsernameBot     string    `json:"username_bot"`
	PasswordBot     string    `json:"password_bot"`
	SelectTextBot   string    `json:"select_text_bot"`
	CreatedBy       int       `gorm:"not null" json:"created_by"`
	UpdatedBy       int       `gorm:"not null" json:"updated_by"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"updated_at"`
}

type AdminSettingSystemResult struct {
	DepositWithdraw bool `json:"deposit_withdraw"`
	Bet             bool `json:"bet"`
	CancelBet       bool `json:"cancel_bet"`
}
