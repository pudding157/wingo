package handlers

import (
	// "encoding/json"
	"fmt"
	"net/http"
	"winapp/internal/app"
	"winapp/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// DB *gorm.DB
	// R  *redis.Client
	c *app.Config
}

func (h *UserHandler) GetProfile(c echo.Context) error {

	auth_header := c.Request().Header.Get("Authorization")
	auth_len := len(auth_header)
	token := auth_header[7:auth_len]

	fmt.Println("token :", token)

	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || t == nil {
		fmt.Println("token err", err)
	}
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	userid := claims["user_id"]
	// userid := c.Param("userid")
	fmt.Println("userid :", userid)
	User := models.User{}
	h.c.DB.Where("id = ?", userid).Find(&User)
	User_bank := models.User_bank{}
	h.c.DB.Where("user_id = ?", User.Id).Find(&User_bank)

	Bank := models.Bank{}
	h.c.DB.Where("id = ?", User_bank.Bank_id).Find(&Bank)

	User_Profile := models.UserProfile{}
	User_Profile.Name = User.First_name + " " + User.Last_name
	User_Profile.PhoneNumber = User.Phone_number
	User_Profile.BankAccount = User_bank.Bank_account
	User_Profile.BankName = Bank.Name

	_res := models.Response{}
	_res.Data = User_Profile // or false
	return c.JSON(http.StatusOK, _res)

	// j, _ := json.Marshal(User_Profile)
	// return c.String(http.StatusOK, string(j))
}
