package middlewares

import (
	"fmt"
	"winapp/app"

	"winapp/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	User_id string `json:"user_id"`
	jwt.StandardClaims
}

// RequestHandlerMiddleware func (Each *Handler)
func RequestHandlerMiddleware(config *app.Config, e *echo.Echo) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			pass := false
			// config := middleware.JWTConfig{
			// 	Claims:     &jwtCustomClaims{},
			// 	SigningKey: []byte("secret"),
			// }
			// var IsLoggedIn = middleware.JWTWithConfig(config)
			// if IsLoggedIn {

			// }
			// m := new(models.User)
			// if err := c.Bind(m); err != nil {
			// 	return utils.JSONResponse(c, nil, nil)
			// }

			// fmt.Println("m out => ", m)
			// for _, mdb := range config.Db.User {
			// 	fmt.Println("m => ", m)
			// 	fmt.Println("mdb => ", mdb)

			// 	if m.Username == mdb.Username && m.Password == mdb.Password {
			// 		pass = true
			// 	}
			// }
			if pass {
				return next(c)
			} else {
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
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
