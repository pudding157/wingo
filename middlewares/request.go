package middlewares

import (
	"encoding/json"
	"fmt"
	"time"

	"winapp/app"
	"winapp/utils"

	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	User_id string `json:"user_id"`
	jwt.StandardClaims
}

type redisValue struct {
	User_id     int    `json:"user_id"`
	Expire_date string `json:"expire_date"`
}

// RequestHandlerMiddleware func (Each *Handler)
func RequestHandlerMiddleware(config *app.Config, e *echo.Echo) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// pass := false
			auth_header := c.Request().Header.Get("Authorization")
			auth_len := len(auth_header)
			token := auth_header[7:auth_len]
			fmt.Println(token)

			_data := config.R.Get(token)
			if _data.Err() != nil {
				fmt.Println("not passed")
				fmt.Println("err => ,,", _data.Err())
				return _data.Err()
			}

			val, _err := _data.Result()
			fmt.Println("_data.Result", val, _err)

			aa := []byte(val)
			redisValue := redisValue{}
			json.Unmarshal(aa, &redisValue)
			fmt.Println("passed", redisValue)

			t, _ := time.Parse(time.RFC3339, redisValue.Expire_date)
			fmt.Println(t)

			now := time.Now()
			diff := t.Sub(now)
			fmt.Printf("Lifespan is %+v", diff)
			_diff := int(diff)
			if _diff <= 0 {
				fmt.Printf("timeup")
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
			} else {
				fmt.Printf("have time")
				return next(c)
			}
		}
	}
}

// RequestMiddleware func (Global)
func RequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		fmt.Println("Requested in global middleware")
		return next(c)
	}
}
