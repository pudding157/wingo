package models

// Bank struct
type User_Bank struct {
	BankId      int    `json:"bank_id"`
	Bank        Bank   `gorm:"foreignKey:Bank_id"`
	BankAccount string `gorm:"unique;not null" json:"bank_account"`
	UserId      int    `json:"user_id"`
	User        User   `gorm:"foreignKey:user_id"`
	Created_at  string `json:"created_at"`
}
