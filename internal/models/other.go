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
	Recipient string `json:"recipient"`
	Type      string `json:"type"`
	Success   bool   `json:"success"`
}

// otp formvalue struct
type OtpResut struct {
	Otp     int  `json:"otp" form:"otp"`
	Success bool `json:"success"`
}

type LoadMoreModel struct {
	Type  string `json:"type"`
	Skip  int    `json:"skip"`
	Take  int    `json:"take"`
	Count int64  `json:"count"`
}
