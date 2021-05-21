package main

import (
	"flag"
	"fmt"

	"winapp/app"
	"winapp/handlers"
	"winapp/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//handler "winapp/handler" //คล้าย namespace c#  ใส่ชื่อapp ดูจาก go.mod ได้หากลืม ช่อง module แล้วใส่ / ชื่อ package
func main() {
	env := flag.String("env", "dev", "environment")
	flag.Parse()
	c := app.NewConfig(*env)
	if err := c.Init(); err != nil {
		fmt.Println(err)
	}
	e := echo.New()

	if err := handlers.NewRouter(e, c); err != nil {
		fmt.Println("New Router Failed.")
	}

	e.Use(middleware.Logger())
	e.Use(middlewares.RequestMiddleware)
	// config := middleware.JWTConfig{
	// 	TokenLookup: "query:token",
	// 	SigningKey:  []byte("secret"),
	// }
	// e.Use(middleware.JWTWithConfig(config))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Logger.Fatal(e.Start(":8000"))
}
