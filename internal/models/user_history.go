package models

import "time"

type User_Bind_History struct {
	AdminBankAccount string  `json:"bank_account"`
	Amount           float64 `json:"amount"`
	Type             int     `json:"type"`
	TransferredAt    string  `json:"transferred_at"`
	Status           int     `json:"status"`
}

// transfers
type User_History struct {
	Id               int       `gorm:"primary_key" json:"id"`
	UserId           int       `json:"user_id"`
	User             User      `json:"-"`
	AdminBankAccount string    `json:"admin_bank_account"`
	Amount           float64   `gorm:"not null;default:0" json:"amount"`
	Type             int       `gorm:"not null;default:0" json:"type"` // withdraw & deposit
	TransferredAt    time.Time `json:"transferred_at"`
	CreatedAt        time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt        time.Time `gorm:"not null" json:"updated_at"`
	Status           int       `gorm:"not null;default:0" json:"status"`
}
