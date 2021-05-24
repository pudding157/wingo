package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// RequestMiddleware func (Global)
func RequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		fmt.Println("Requested in global middleware")
		return next(c)
	}
}
