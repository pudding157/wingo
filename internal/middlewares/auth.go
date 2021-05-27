package middlewares

import (
	"encoding/json"
	"fmt"
	"time"
	"winapp/internal/app"
	"winapp/internal/utils"

	"github.com/labstack/echo/v4"
)

//
type UserSession struct {
	UserID     int    `json:"user_id"`
	ExpireDate string `json:"expire_date"`
}

// AuthMiddleware func (Each *Handler)
func AuthMiddleware(config *app.Config, e *echo.Echo) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// pass := false
			auth_header := c.Request().Header.Get("Authorization")
			auth_len := len(auth_header)
			token := auth_header[7:auth_len]
			fmt.Println(token)
			d := config.R.Get(token)
			if d.Err() != nil {
				fmt.Println("not passed")
				fmt.Println("err => ,,", d.Err())
				return d.Err()
			}

			val, _err := d.Result()
			fmt.Println("_data.Result", val, _err)

			aa := []byte(val)
			rv := UserSession{}
			json.Unmarshal(aa, &rv)
			fmt.Println("passed", rv)

			t, _ := time.Parse(time.RFC3339, rv.ExpireDate)
			fmt.Println(t)

			now := time.Now()
			diff := t.Sub(now)
			fmt.Printf("Lifespan is %+v \n", diff)
			_diff := int(diff)

			config.T = ""
			config.UI = 0
			if _diff <= 0 {
				fmt.Printf("timeup")
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
			} else {
				fmt.Printf("have time")

				config.T = token
				config.UI = rv.UserID
				return next(c)
			}
		}
	}
}
