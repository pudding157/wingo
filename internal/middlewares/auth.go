package middlewares

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"winapp/internal/app"
	"winapp/internal/utils"

	"github.com/labstack/echo/v4"
)

//
type UserSession struct {
	UserID     int    `json:"user_id"`
	ExpireDate string `json:"expire_date"`
	RoleName   string `json:"role_name"`
}

// will check role for endpoint
func AuthMiddleware(cf *app.Config, e *echo.Echo) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// pass := false
			auth_header := c.Request().Header.Get("Authorization")
			fmt.Println("auth_header = ", auth_header)
			hb := strings.Contains(auth_header, "Bearer ")
			if !hb {
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
			}
			auth_len := len(auth_header)
			token := auth_header[7:auth_len]
			fmt.Println(token)
			d := cf.R.Get(token)
			if d.Err() != nil {
				fmt.Println("not passed")
				fmt.Println("err => ,,", d.Err())
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
			}

			val, _err := d.Result()
			fmt.Println("_data.Result", val, _err)

			aa := []byte(val)
			rv := UserSession{}
			json.Unmarshal(aa, &rv)
			fmt.Println("passed", rv)

			t, _ := time.Parse(time.RFC3339, rv.ExpireDate)
			fmt.Println(t)

			now := time.Now().UTC()
			diff := t.Sub(now)
			fmt.Printf("Lifespan is %+v \n", diff)
			_diff := int(diff)

			cf.T = ""
			cf.UI = 0
			if _diff <= 0 {
				fmt.Println("timeup")
				return utils.JSONResponse(c, nil, utils.NewUnauthorizedError())
			} else {
				fmt.Println("have time")

				cf.T = token
				cf.UI = rv.UserID
				cf.ROLE = rv.RoleName

				fmt.Println("cf.ROLE => ", cf.ROLE)
				return next(c)
			}
		}
	}
}
