package models

// เดี๋ยวลบ
// Merchant struct
type Merchant struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	BangkAccount string `json:"bankAccount,omitempty"`
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
}
