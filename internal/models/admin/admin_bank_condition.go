package models

import (
	"time"
	"winapp/internal/models"
)

// User struct
type Admin_Bank_Condition struct {
	Id          int         `gorm:"primary_key" json:"id"`
	BankId      int         `json:"bank_id"`
	Bank        models.Bank `gorm:"foreignKey:bank_id" json:"-"`
	BankAccount string      `gorm:"not null" json:"bank_account"`
	PriceStart  float64     `gorm:"not null;default:0" json:"price_start"`
	PriceEnd    float64     `gorm:"not null;default:0" json:"price_end"`
	IsActive    bool        `gorm:"not null;default:1" json:"is_active"`
	DeviceId    string      `json:"device_id"`
	ApiRefresh  string      `json:"api_refresh"`
	AccessToken string      `json:"accesstoken"`
	CreatedAt   time.Time   `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"not null" json:"updated_at"`
}
