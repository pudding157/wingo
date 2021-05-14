package controller

import (
	"fmt"
	"winapp/util"

	"net/http"

	"github.com/labstack/echo/v4"
)

// register action form
func Register(c echo.Context) error {

	fmt.Println("Register Controller")

	_str := "myPassword"

	_hashStr := util.HashStr(_str)

	fmt.Println("string is => ", _str)
	fmt.Println("Hash is => ", _hashStr)

	//test dehash
	isMatch := util.DehashStr(_hashStr, _str)

	fmt.Println("isMatch ? => ", isMatch)

	return c.JSON(http.StatusOK, _hashStr)
}
