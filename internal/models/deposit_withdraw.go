package models

type Deposit_Information struct {
	Id          int     `gorm:"primary_key" json:"id"`
	BankAccount string  `gorm:"not null" json:"bank_account"`
	MoneyAmount float64 `gorm:"not null;default:0" json:"money_amount"`
	TransferAt  string  `json:"transfer_at"`
	Updated_at  string  `json:"updated_at"`
	Status      int     `gorm:"not null;default:0" json:"status"`
}

type Withdraw_Information struct {
	Id int `gorm:"primary_key" json:"id"`
}
