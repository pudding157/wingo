package handlers

import (
	// "encoding/json"

	"fmt"
	"net/http"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.PaymentRepository
}

func NewPaymentHandler(repo repositories.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{Repo: repo}
}

func (r *PaymentHandler) Deposit(c echo.Context) error {
	_res := models.Response{}
	ub := models.User_bind_history{}
	c.Bind(ub)
	uh := models.User_History{}
	uh.AdminBankAccount = ub.AdminBankAccount
	uh.TransferredAt = ub.TransferredAt
	uh.Amount = ub.Amount
	uh.Status = ub.Status

	err := r.Repo.Deposit(uh)
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = true
	return c.JSON(http.StatusOK, _res)
}

func (r *PaymentHandler) Withdraw(c echo.Context) error {
	_res := models.Response{}

	err := r.Repo.Withdraw()
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = "withdraw"
	return c.JSON(http.StatusOK, _res)
}

func (r *PaymentHandler) Transactions(c echo.Context) error {
	_res := models.Response{}

	t := c.Param("type")
	fmt.Println("param => ", t)

	err := r.Repo.Transactions(t)
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = "withdraw"
	return c.JSON(http.StatusOK, _res)
}
