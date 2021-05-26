package models

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	User_id string `json:"user_id"`
	jwt.StandardClaims
}

type RedisValue struct {
	User_id     int    `json:"user_id"`
	Expire_date string `json:"expire_date"`
}

// otp formvalue struct
type OtpModel struct {
	Otp       int    `json:"otp" form:"otp"`
	Recipient string `json:"-" form:"recipient"`
	Type      string `json:"-" form:"type"`
	Success   bool   `json:"success"`
}
