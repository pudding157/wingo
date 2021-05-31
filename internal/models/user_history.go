package models

type User_bind_history struct {
	AdminBankAccount string  `json:"bank_account"`
	TransferredAt    string  `json:"transferred_at"`
	Amount           float64 `json:"amount"`
	Status           int     `json:"status"`
}

// transfers
type User_History struct {
	Id               int `gorm:"primary_key" json:"id"`
	UserId           int `json:"user_id"`
	User             User
	AdminBankAccount string  `json:"admin_bank_account"`
	Amount           float64 `gorm:"not null;default:0" json:"amount"`
	Type             int     `gorm:"not null;default:0" json:"type"` // withdraw & deposit
	TransferredAt    string  `json:"transferred_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Status           int     `gorm:"not null;default:0" json:"status"`
}
