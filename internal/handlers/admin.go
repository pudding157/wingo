package handlers

import (
	"fmt"
	"strconv"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AdminHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.AdminRepository
}

func NewAdminHandler(repo repositories.AdminRepository) *AdminHandler {
	return &AdminHandler{Repo: repo}
}

func (r *AdminHandler) PostHome(c echo.Context) error {

	fmt.Println("Post all home details")

	pc := &models.Page_Content{}
	c.Bind(&pc)
	h, err := r.Repo.PostHome(*pc)

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = h

	return c.JSON(http.StatusOK, _res)
}

func (r *AdminHandler) PostBlog(c echo.Context) error {

	fmt.Println("Post all home details")

	bc := &models.Blog_Content{}
	c.Bind(&bc)

	fmt.Println("Blog_Content => ", bc)

	h, err := r.Repo.PostBlog(*bc)
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = h

	return c.JSON(http.StatusOK, _res)
}

func (r *AdminHandler) GetWallets(c echo.Context) error {

	fmt.Println("Get all user Wallets")

	w, err := r.Repo.GetWallets()
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = w

	return c.JSON(http.StatusOK, _res)
}
