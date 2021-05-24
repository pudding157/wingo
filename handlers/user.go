package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"winapp/app"
	"winapp/models"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	// DB *gorm.DB
	// R  *redis.Client
	c *app.Config
}

// func UserHandler(c *app.Config) *UserHandler {

// 	return &UserHandler{DB: c.DB, R: c.R}
// }

func (h *UserHandler) Get_Profile(c echo.Context) error {
	userid := c.Param("userid")
	fmt.Println("userid :", userid)
	User := models.User{}
	h.c.DB.Where("id = ?", userid).Find(&User)
	j, _ := json.Marshal(User)
	return c.String(http.StatusOK, string(j))
}
