package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//handler "winapp/handler" //คล้าย namespace c#  ใส่ชื่อapp ดูจาก go.mod ได้หากลืม ช่อง module แล้วใส่ / ชื่อ package
func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	//กำหนด Route ก่อนเลย พร้อมให้ call ไปยัง func ต่างๆ
	// userHandler := handler.UserHandler{}
	// userHandler.Initialize() //เชื่อมต่อเมื่อเริ่ม

	// redisHandler := handler.RedisHandler{}
	// redisHandler.Initialize()

	register_module(e)
	login_module(e)

	e.Logger.Fatal(e.Start(":8000"))
}

func register_module(e *echo.Echo) {

	// get all bank
	e.GET("/api/v1/bank", func(c echo.Context) error {
		return c.String(http.StatusOK, "get all bank")
	})

	// send otp
	e.POST("/api/v1/register/otp/send", func(c echo.Context) error {
		return c.String(http.StatusOK, "send otp")
	})

	// post check otp
	e.POST("/api/v1/register/otp", func(c echo.Context) error {
		return c.String(http.StatusOK, "post check otp")
	})

	// register form submit
	e.POST("/api/v1/register", func(c echo.Context) error {
		return c.String(http.StatusOK, "register form submit")
	})
}

func login_module(e *echo.Echo) {

	// login form submit
	e.POST("/api/v1/login", func(c echo.Context) error {
		return c.String(http.StatusOK, "login form submit")
	})

}
