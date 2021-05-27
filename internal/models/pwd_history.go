package models

type Password_History struct {
	Id          int    `gorm:"primary_key" json:"id"`
	Username    string `gorm:"not null" json:"username"`
	IPAddress   string `json:"ip_address"`
	MACAddress  string `json:"mac_address"`
	Browser     string `json:"browser"`
	Created_at  string `json:"created_at"`
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}
