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

type BankHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.BankRepository
}

func NewBankHandler(repo repositories.BankRepository) *BankHandler {
	return &BankHandler{Repo: repo}
}

// otp/send action form
func (r *BankHandler) GetBanks(c echo.Context) error {

	fmt.Println("Get all bank")

	banks, err := r.Repo.GetBanks()

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = banks

	return c.JSON(http.StatusOK, _res)
}

// otp/send action form
func (r *BankHandler) GetAdminBanks(c echo.Context) error {

	fmt.Println("Get all bank")

	banks, err := r.Repo.GetAdminBanks()

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = banks

	return c.JSON(http.StatusOK, _res)
}
