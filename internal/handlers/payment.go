package handlers

import (
	// "encoding/json"

	"fmt"
	"net/http"
	"strconv"
	"time"
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
	ub := models.User_Bind_History{}
	c.Bind(&ub)
	uh := models.User_History{}
	uh.AdminBankAccount = ub.AdminBankAccount
	// layout := "Tue, 01 Jun 2021 01:21:00 GMT"
	// ta, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	ta, err := time.Parse(time.RFC1123, ub.TransferredAt)
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}
	uh.TransferredAt = ta
	fmt.Println("uh.TransferredAt", uh.TransferredAt)
	uh.Amount = ub.Amount
	uh.Status = ub.Status

	fmt.Println("uh => ", uh)

	err = r.Repo.Deposit(uh)
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = map[string]bool{
		"success": true,
	}
	return c.JSON(http.StatusOK, _res)
}

func (r *PaymentHandler) Withdraw(c echo.Context) error {
	_res := models.Response{}
	ub := models.User_Bind_History{} // แค่ไว้รับ
	c.Bind(&ub)
	uh := models.User_History{}
	uh.Amount = ub.Amount

	err := r.Repo.Withdraw(uh)
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = map[string]bool{
		"success": true,
	}
	return c.JSON(http.StatusOK, _res)
}

func (r *PaymentHandler) Transactions(c echo.Context) error {
	_res := models.Response{}

	qs := c.QueryParams()
	fmt.Println("QueryParams => ", qs)

	l := models.LoadMoreModel{}

	l.Type = qs.Get("type")

	if l.Type != "all" && l.Type != "deposit" && l.Type != "withdraw" {
		l.Type = "all"
	}
	s, err := strconv.Atoi(qs.Get("skip"))
	if err != nil {
		s = 0
	}
	l.Skip = s

	t, err := strconv.Atoi(qs.Get("take"))
	if err != nil {
		t = 20
	}
	l.Take = t

	tr, cnt, err := r.Repo.Transactions(l)

	l.Count = cnt
	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}
	_res.Data = map[string]interface{}{
		"data":    tr,
		"filters": l,
	} // or false
	// _res.Data = tr
	return c.JSON(http.StatusOK, _res)
}
