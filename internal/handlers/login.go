package handlers

import (
	"fmt"
	"net/http"
	"strconv"
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
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
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
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = map[string]bool{
		"success": true,
	}

	return c.JSON(http.StatusOK, _res)
	// return c.JSON(http.StatusOK, map[string]string{
	// 	"token": t,
	// })
}
