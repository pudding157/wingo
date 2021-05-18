package handlers

import (
	"fmt"
	"log"
	"winapp/models"
	"winapp/utils"

	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Handler struct
type Handler struct {
	DB *gorm.DB
	// mService services.MerchantService
	// pService services.ProductService
	// rService services.ReportService
}

func RegisterHandler() *Handler {
	return &Handler{}
}

// otp formvalue struct
type RegisterFormModel struct {
	First_name   string
	Last_name    string
	Phone_number string
	Bank_id      int
	Bank_account string
	Username     string
	Password     string
}
type ReturnToken struct {
	Token string `json:"token"`
}

// register action form
func (h Handler) Register(c echo.Context) error {

	fmt.Println("Register")

	registerFormModel := RegisterFormModel{}
	registerFormModel.First_name = c.FormValue("first_name")     // get params
	registerFormModel.Last_name = c.FormValue("last_name")       // get params
	registerFormModel.Phone_number = c.FormValue("phone_number") // get params
	_bank_id, err := strconv.Atoi(c.FormValue("bank_id"))
	if err != nil {
		log.Fatal(err)
	}
	registerFormModel.Bank_id = _bank_id                         // get params
	registerFormModel.Bank_account = c.FormValue("bank_account") // get params
	registerFormModel.Username = c.FormValue("username")         // get params
	registerFormModel.Password = c.FormValue("password")         // get params

	fmt.Println("registerFormModel => ", registerFormModel)

	_passwordHashed := utils.HashStr(registerFormModel.Password)

	fmt.Println("Hash is => ", _passwordHashed)

	//test dehash
	// isMatch := utils.DehashStr(_passwordHashed, _str)

	// fmt.Println("isMatch ? => ", isMatch)

	_res := models.Response{}

	_res.Data = ReturnToken{Token: "asd"}

	return c.JSON(http.StatusOK, _res)
}

// otp/send action form
func (h Handler) Otp_send(c echo.Context) error {

	fmt.Println("Otp send")

	phone_number := c.FormValue("phone_number") // get params
	fmt.Println("phone_number => ", phone_number)
	_res := models.Response{}
	_res.Message = "true" // or false

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.JSON(http.StatusOK, _res)
}

// otp formvalue struct
type OtpModel struct {
	Otp       string
	Recipient string
	Type      string
}

// otp action form
func (h Handler) Otp(c echo.Context) error {

	fmt.Println("Otp")
	otpModel := OtpModel{}
	otpModel.Otp = c.FormValue("otp")             // get params
	otpModel.Recipient = c.FormValue("recipient") // get params
	otpModel.Type = c.FormValue("type")           // get params

	fmt.Println("model => ", otpModel)

	_res := models.Response{}
	_res.Message = "true" // or false

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.JSON(http.StatusOK, _res)
}
