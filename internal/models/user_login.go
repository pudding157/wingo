package models

import "time"

// log_login struct
type User_Login struct {
	Token      string    `gorm:"primary_key" json:"token"`
	Id         int       `gorm:"default:1;uniqueIndex;autoIncrement;not null" json:"id"`
	UserId     int       `json:"user_id"`
	User       User      `gorm:"foreignKey:User_id"`
	Username   string    `json:"username"`
	IpAddress  string    ` json:"ip_address"`
	MacAddress string    ` json:"mac_address"`
	UserAgent  string    `json:"user_agent"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
}
