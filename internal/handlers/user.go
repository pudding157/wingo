package handlers

import (
	// "encoding/json"

	"fmt"
	"net/http"
	"strconv"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"github.com/labstack/echo/v4"
)

// type UserHandler struct {
// 	// DB *gorm.DB
// 	// R  *redis.Client
// 	c *app.Config
// }

type UserHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (r *UserHandler) GetProfile(c echo.Context) error {

	// auth_header := c.Request().Header.Get("Authorization")
	// auth_len := len(auth_header)
	// token := auth_header[7:auth_len]

	up, err := r.Repo.GetProfile()

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = up // or false
	return c.JSON(http.StatusOK, _res)

	// j, _ := json.Marshal(User_Profile)
	// return c.String(http.StatusOK, string(j))
}

func (r *UserHandler) ChangePassword(c echo.Context) error {

	ph := models.Password_History{}
	c.Bind(&ph)

	t, err := r.Repo.ChangePassword(ph)
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
	} // or false
	return c.JSON(http.StatusOK, _res)
}

func (r *UserHandler) GetAffiliate(c echo.Context) error {
	_res := models.Response{}

	fmt.Println("param => ")

	a, err := r.Repo.GetAffiliate()
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = a
	return c.JSON(http.StatusOK, _res)
}
