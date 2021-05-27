package models

import "github.com/dgrijalva/jwt-go"

type JWTCustomClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

type RedisValue struct {
	UserId     int    `json:"user_id"`
	ExpireDate string `json:"expire_date"`
}

// otp formvalue struct
type OtpModel struct {
	Otp       int    `json:"otp" form:"otp"`
	Recipient string `json:"-" form:"recipient"`
	Type      string `json:"-" form:"type"`
	Success   bool   `json:"success"`
}
