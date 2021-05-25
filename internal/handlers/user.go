package handlers

import (
	// "encoding/json"

	"net/http"
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

	auth_header := c.Request().Header.Get("Authorization")
	auth_len := len(auth_header)
	token := auth_header[7:auth_len]

	up, err := r.Repo.GetProfile(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	_res := models.Response{}
	_res.Data = up // or false
	return c.JSON(http.StatusOK, _res)

	// j, _ := json.Marshal(User_Profile)
	// return c.String(http.StatusOK, string(j))
}
