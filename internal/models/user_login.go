package models

import "time"

// log_login struct
type User_Login struct {
	Token       string    `gorm:"primary_key" json:"token"`
	Id          int       `gorm:"default:1;uniqueIndex;autoIncrement;not null" json:"id"`
	User_id     int       `json:"user_id"`
	User        User      `gorm:"foreignKey:User_id"`
	Username    string    `json:"username"`
	Ip_address  string    ` json:"ip_address"`
	Mac_address string    ` json:"mac_address"`
	User_agent  string    `json:"user_agent"`
	Created_at  time.Time `gorm:"not null" json:"created_at"`
}
