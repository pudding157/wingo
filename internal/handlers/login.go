package handlers

import (
	"fmt"
	"net/http"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.LoginRepository
}

func NewLoginHandler(repo repositories.LoginRepository) *LoginHandler {
	return &LoginHandler{Repo: repo}
}

func (r *LoginHandler) Login(c echo.Context) error {

	bu := &models.User{}
	c.Bind(&bu)
	t, err := r.Repo.Login(*bu)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_res := models.Response{}
	_res.Data = map[string]string{
		"token": *t,
	}

	return c.JSON(http.StatusOK, _res)
	// return c.JSON(http.StatusOK, map[string]string{
	// 	"token": t,
	// })
}

func (r *LoginHandler) Logout(c echo.Context) error {

	fmt.Println("Logout")
	err := r.Repo.Logout()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_res := models.Response{}
	_res.Data = true

	return c.JSON(http.StatusOK, _res)
	// return c.JSON(http.StatusOK, map[string]string{
	// 	"token": t,
	// })
}
