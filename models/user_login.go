package models

// log_login struct
type User_login struct {
	Token       string `gorm:"primary_key" json:"token"`
	Id          int    `gorm:"type:autoIncrement" json:"id"`
	User_id     int    `json:"user_id"`
	User        User   `gorm:"foreignKey:User_id"`
	Username    string `gorm:"not null" json:"username"`
	Ip_address  string `gorm:"not null" json:"ip_address"`
	Mac_address string `gorm:"not null" json:"mac_address"`
	User_agent  string `gorm:"not null" json:"user_agent"`
	Created_at  string `json:"created_at"`
}
