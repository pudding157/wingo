package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"winapp/internal/app"
	"winapp/internal/models"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// DB *gorm.DB
	// R  *redis.Client
	c *app.Config
}

type User_Profile struct {
	Name         string `json:"name"`
	Phone_number string `json:"phone_number"`
	Bank_name    string `json:"bank_name"`
	Bank_account string `json:"bank_account"`
	Status       string `json:"status"`
}

func (h *UserHandler) Get_Profile(c echo.Context) error {
	userid := c.Param("userid")
	fmt.Println("userid :", userid)
	User := models.User{}
	h.c.DB.Where("id = ?", userid).Find(&User)
	User_bank := models.User_bank{}
	h.c.DB.Where("user_id = ?", User.Id).Find(&User_bank)

	User_Profile := User_Profile{}
	User_Profile.Name = User.First_name + " " + User.Last_name

	// User_Profile.Bank_name = User.Bank.Name
	// User_Profile.Bank_account = User

	j, _ := json.Marshal(User_Profile)
	return c.String(http.StatusOK, string(j))
}
