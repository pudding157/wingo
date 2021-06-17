package models

// Bank struct
type Bank struct {
	Id       int    `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	IsActive bool   `gorm:"default:false;column:is_active" json:"-"`
}

type AdminBank struct {
	Id          int    `gorm:"primary_key;autoIncrement" json:"id"`
	BankId      int    `json:"bank_id"`
	Bank        Bank   `gorm:"foreignKey:bank_id" json:"-"`
	BankAccount string `gorm:"unique;not null" json:"bank_account"`
	IsActive    bool   `gorm:"default:false" json:"-"`
}

type AdminBankModel struct {
	BankId      int    `json:"bank_id"`
	BankName    string `json:"name"`
	BankAccount string `gorm:"unique;not null" json:"bank_account"`
}
