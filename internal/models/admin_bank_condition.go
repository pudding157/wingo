package models

import (
	"time"
)

// Admin_Bank_Condition struct
type Admin_Bank_Condition struct {
	Id          int       `gorm:"primary_key" json:"id"`
	BankId      int       `json:"bank_id"`
	Bank        Bank      `gorm:"foreignKey:bank_id" json:"-"`
	BankAccount string    `json:"bank_account"`
	PriceStart  float64   `gorm:"not null;default:0" json:"price_start"`
	PriceEnd    float64   `gorm:"not null;default:0" json:"price_end"`
	IsActive    bool      `gorm:"not null;default:1" json:"is_active"`
	DeviceId    string    `json:"device_id"`
	ApiRefresh  string    `json:"api_refresh"`
	AccessToken string    `json:"accesstoken"`
	CreatedBy   int       `gorm:"not null" json:"created_by"`
	UpdatedBy   int       `gorm:"not null" json:"updated_by"`
	CreatedAt   time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time `gorm:"not null" json:"updated_at"`
}

type AdminSettingBotListResult struct {
	IsBotActive           bool                    `json:"is_bot_active"`
	AdminSettingBotResult []AdminSettingBotResult `json:"condition_list"`
}

type AdminSettingBotResult struct {
	Id          int     `json:"id"`
	IsActive    bool    `json:"is_active"`
	PriceStart  float64 `json:"price_start"`
	PriceEnd    float64 `json:"price_end"`
	BankId      int     `json:"bank_id"`
	BankAccount string  `json:"bank_account"`
}
