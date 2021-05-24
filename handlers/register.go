package handlers

import (
	"fmt"
	"log"
	"time"
	"winapp/app"
	"winapp/enums"
	"winapp/models"
	"winapp/utils"

	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RegisterHandler struct {
	// DB *gorm.DB
	// R  *redis.Client
	c *app.Config
}

// func RegisterHandler(c *app.Config) *Handler {
// 	Otp_history := []models.Otp_history{}
// 	if !c.DB.HasTable(Otp_history) {
// 		fmt.Println("No table")
// 		c.DB.AutoMigrate(&Otp_history) // สร้าง table, field ต่างๆที่ไม่เคยมี
// 		fmt.Println("migrate data Otp_history")
// 	}
// 	User := []models.User{}
// 	if !c.DB.HasTable(User) {
// 		fmt.Println("No table")
// 		c.DB.AutoMigrate(&User) // สร้าง table, field ต่างๆที่ไม่เคยมี
// 		fmt.Println("migrate data User")
// 	}

// 	User_bank := models.User_bank{}
// 	if !c.DB.HasTable(User_bank) {
// 		fmt.Println("No table")
// 		c.DB.AutoMigrate(&User_bank) // สร้าง table, field ต่างๆที่ไม่เคยมี
// 		fmt.Println("migrate data User_bank")
// 	} else {
// 		// ถ้าจะเพิ่ม unique ถ้า ในตารางมีข้อมูลซ้ำจะไม่สามารถทำได้
// 		c.DB.Model(&User_bank).AddUniqueIndex("bank_account", "bank_account")
// 	}

// 	return &Handler{DB: c.DB, R: c.R}
// }

// otp formvalue struct
type RegisterFormModel struct {
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Phone_number string `json:"phone_number"`
	Bank_id      int    `json:"bank_id"`
	Bank_account string `json:"bank_account"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Otp          int    `json:"otp"`
}
type ReturnToken struct {
	Token string `json:"token"`
}

// register action form
func (h *RegisterHandler) Register(c echo.Context) error {

	fmt.Println("Register")

	Bind_registerFormModel := &RegisterFormModel{}
	c.Bind(&Bind_registerFormModel)

	fmt.Println("Bind_registerFormModel, ", Bind_registerFormModel)

	// registerFormModel := RegisterFormModel{}
	// registerFormModel.First_name = c.FormValue("first_name")     // get params
	// registerFormModel.Last_name = c.FormValue("last_name")       // get params
	// registerFormModel.Phone_number = c.FormValue("phone_number") // get params
	// _bank_id, err := strconv.Atoi(c.FormValue("bank_id"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// registerFormModel.Bank_id = _bank_id                         // get params
	// registerFormModel.Bank_account = c.FormValue("bank_account") // get params
	// registerFormModel.Username = c.FormValue("username")         // get params
	// registerFormModel.Password = c.FormValue("password")         // get params
	// registerFormModel.Otp = c.FormValue("otp")                   // get params

	// fmt.Println("registerFormModel => ", registerFormModel)

	_passwordHashed := utils.HashStr(Bind_registerFormModel.Password)

	fmt.Println("Hash is => ", _passwordHashed)

	User := models.User{}

	User.First_name = Bind_registerFormModel.First_name
	User.Last_name = Bind_registerFormModel.Last_name
	User.Phone_number = Bind_registerFormModel.Phone_number
	User.Username = Bind_registerFormModel.Username
	User.Password = _passwordHashed
	_now := time.Now().Format(time.RFC3339)
	User.Created_at = _now
	User.Updated_at = _now
	User.Registration_otp = strconv.Itoa(Bind_registerFormModel.Otp)

	if err := h.c.DB.Save(&User).Error; err != nil {
		log.Print("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_message = err
		_res.Error_code = "400"
		return c.JSON(http.StatusBadRequest, _res)
	}

	User_bank := models.User_bank{}
	User_bank.Bank_id = Bind_registerFormModel.Bank_id
	User_bank.Bank_account = Bind_registerFormModel.Bank_account
	User_bank.User_id = User.Id
	User_bank.Created_at = _now

	if err := h.c.DB.Save(&User_bank).Error; err != nil {
		log.Print("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_message = err
		_res.Error_code = "400"
		return c.JSON(http.StatusBadRequest, _res)
	}

	/*
			type User_login struct {
			Token       string `gorm:"primary_key" json:"token"`
			Id          int    `gorm:"type:autoIncrement" json:"id"`
			User_id     int    `json:"user_id"`
			User        User   `gorm:"foreignKey:User_id"`
			Username    string `gorm:"not null" json:"username"`
			Ip_address  string `gorm:"not null" json:"ip_address"`
			Mac_address string `gorm:"not null" json:"mac_address"`
			User_agent  string `gorm:"not null" json:"user_agent"`
			Created_at  string `json:"created_at"`
		}
	*/

	_res := models.Response{}

	_res.Data = ReturnToken{Token: "asd"} // will change

	return c.JSON(http.StatusOK, _res)
}

// otp/send action form
func (h *RegisterHandler) Otp_send(c echo.Context) error {

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

	if err := h.c.DB.Save(&history).Error; err != nil {
		log.Print("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "400"
		return c.JSON(http.StatusBadRequest, _res)
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
func (h *RegisterHandler) Otp(c echo.Context) error {

	fmt.Println("Otp")
	otpModel := OtpModel{}

	_otp, _ := strconv.Atoi(c.FormValue("otp"))
	otpModel.Otp = _otp                           // get params
	otpModel.Recipient = c.FormValue("recipient") // get params
	otpModel.Type = c.FormValue("type")           // get params
	keyType, _err := enums.EnumFromKey(otpModel.Type)
	if _err != nil {
		log.Fatal(_err)
	}
	history := models.Otp_history{}
	h.c.DB.Where("otp = ? AND type = ? AND send_to = ?", otpModel.Otp, keyType.Index(), otpModel.Recipient).Find(&history)

	if history.Id == 0 {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "400"
		return c.JSON(http.StatusBadRequest, _res)
	}

	fmt.Println("model => ", history)

	_res := models.Response{}
	otp_res := ReturnOtp{Success: true, Otp: otpModel.Otp}
	otp_res.Success = true
	_res.Data = otp_res // or false

	return c.JSON(http.StatusOK, _res)
}
