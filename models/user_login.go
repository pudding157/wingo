package models

// log_login struct
type User_login struct {
	Token       string `gorm:"primary_key" json:"token"`
	Id          int    `gorm:"autoIncrement" json:"id"`
	User_id     int    `json:"user_id"`
	User        User   `gorm:"foreignKey:User_id"`
	Username    string `json:"username"`
	Ip_address  string ` json:"ip_address"`
	Mac_address string ` json:"mac_address"`
	User_agent  string `json:"user_agent"`
	Created_at  string `json:"created_at"`
}
