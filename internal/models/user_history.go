package models

// transfers
type User_History struct {
	Id               int     `gorm:"primary_key" json:"id"`
	UserId           int     `json:"user_id"`
	User             User    `gorm:"foreignKey:user_id"`
	AdminBankAccount string  `json:"admin_bank_account"`
	Amount           float64 `gorm:"not null;default:0" json:"amount"`
	Type             int     `gorm:"not null;default:0" json:"type"` // withdraw & deposit
	TransferredAt    string  `json:"transferred_at"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	Status           int     `gorm:"not null;default:0" json:"status"`
}
