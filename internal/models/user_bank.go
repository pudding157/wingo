package models

// Bank struct
type User_Bank struct {
	Bank_id      int    `json:"bank_id"`
	Bank         Bank   `gorm:"foreignKey:Bank_id"`
	Bank_account string `gorm:"unique;not null" json:"bank_account"`
	User_id      int    `json:"user_id"`
	User         User   `gorm:"foreignKey:User_id"`
	Created_at   string `json:"created_at"`
}

type Admin_Bank struct {
	Id           int    `gorm:"primary_key;autoIncrement" json:"id"`
	BankId       int    `json:"bank_id"`
	Bank         Bank   `gorm:"foreignKey:bank_id"`
	Bank_account string `gorm:"unique;not null" json:"bank_account"`
	UserId       int    `json:"user_id"`
	Created_at   string `json:"created_at"`
}
