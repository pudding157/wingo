package models

import "time"

// User_Wallet struct
type User_Wallet struct {
	Id        int       `gorm:"primary_key" json:"id"`
	UserId    int       `json:"user_id"`
	User      User      `gorm:"foreignKey:User_id"`
	Amount    float64   `gorm:"not null;default:0" json:"amount"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type WalletsResult struct {
	Amount float64 `json:"amount"`
	Count  int     `json:"count"`
}
