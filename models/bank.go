package models

// Bank struct
type Bank struct {
	Id        int    `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	Is_active string `json:"is_active"`
}
