package handlers

import (
	"fmt"
	"winapp/utils"

	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler struct
type Handler struct {
	// mService services.MerchantService
	// pService services.ProductService
	// rService services.ReportService
}

func RegisterHandler() *Handler {
	return &Handler{}
}

// register action form
func (h Handler) Register(c echo.Context) error {

	fmt.Println("Register Controller")

	_str := "myPassword"

	_hashStr := utils.HashStr(_str)

	fmt.Println("string is => ", _str)
	fmt.Println("Hash is => ", _hashStr)

	//test dehash
	isMatch := utils.DehashStr(_hashStr, _str)

	fmt.Println("isMatch ? => ", isMatch)

	return c.JSON(http.StatusOK, _hashStr)
}
