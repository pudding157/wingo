package handlers

import (
	"fmt"
	"log"
	"time"
	"winapp/enums"
	"winapp/models"
	"winapp/utils"

	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func RegisterHandler(db *gorm.DB) *Handler {
	Otp_history := []models.Otp_history{}
	if !db.HasTable(Otp_history) {
		fmt.Println("No table")
		db.AutoMigrate(&Otp_history) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data bank and create bank")
	}

	return &Handler{DB: db}
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
	// fmt.Println("phone_number => ", phone_number)

	// todo otp
	otp := 123456

	_res := models.Response{}
	otp_res := ReturnOtp{}
	otp_res.Success = true
	otp_res.Otp = otp // get params
	_res.Data = otp_res

	history := models.Otp_history{}
	history.Type = enums.PHONE_NUMBER.Index()
	history.Send_to = phone_number
	history.Otp = otp
	history.Created_at = time.Now().Format(time.RFC3339)
	fmt.Println("history => ", history)

	if err := h.DB.Save(&history).Error; err != nil {
		log.Print("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "400"
		return c.JSON(http.StatusOK, _res)
	}
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.JSON(http.StatusOK, _res)
}

// otp formvalue struct
type OtpModel struct {
	Otp       int
	Recipient string
	Type      string
}
type ReturnOtp struct {
	Success bool `json:"success"`
	Otp     int  `json:"otp"`
}

// otp action form
func (h Handler) Otp(c echo.Context) error {

	fmt.Println("Otp")
	otpModel := OtpModel{}

	_otp, _ := strconv.Atoi(c.FormValue("otp"))
	otpModel.Otp = _otp                           // get params
	otpModel.Recipient = c.FormValue("recipient") // get params
	otpModel.Type = c.FormValue("type")           // get params

	fmt.Println("model => ", otpModel)

	_res := models.Response{}
	otp_res := ReturnOtp{}
	otp_res.Success = true
	_res.Data = otp_res // or false
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.JSON(http.StatusOK, _res)
}
