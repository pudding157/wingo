package models

// Bank struct
type User_bank struct {
	Bank_id      int    `json:"bank_id"`
	Bank         Bank   `gorm:"foreignKey:Bank_id"`
	Bank_account string `gorm:"unique;not null" json:"bank_account"`
	User_id      int    `json:"user_id"`
	User         User   `gorm:"foreignKey:User_id"`
	Created_at   string `json:"created_at"`
}