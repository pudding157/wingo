package handlers

import (
	"fmt"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"net/http"

	"github.com/labstack/echo/v4"
)

// type RegisterHandler struct {
// 	// DB *gorm.DB
// 	// R  *redis.Client
// 	c *app.Config
// }
type RegisterHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.RegisterRepository
}

func NewRegisterHandler(repo repositories.RegisterRepository) *RegisterHandler {
	return &RegisterHandler{Repo: repo}
}

// register action form
func (r *RegisterHandler) Register(c echo.Context) error {

	fmt.Println("Register")

	Bind_registerFormModel := &models.RegisterFormModel{}
	c.Bind(&Bind_registerFormModel)

	fmt.Println("Bind_registerFormModel, ", Bind_registerFormModel)
	t, err := r.Repo.Register(*Bind_registerFormModel)

	if err != nil {
		fmt.Println("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}

	tk := *t
	_res := models.Response{}
	_res.Data = map[string]string{
		"token": tk,
	}

	return c.JSON(http.StatusOK, _res)
}

// otp/send action form
func (r *RegisterHandler) Otp_send(c echo.Context) error {

	fmt.Println("Otp send")
	UserProfile := models.UserProfile{}
	c.Bind(&UserProfile) // get params
	phone_number := UserProfile.PhoneNumber
	fmt.Println("phone_number => ", phone_number)

	om, err := r.Repo.Otp_send(phone_number)
	if err != nil {
		fmt.Println("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}

	_res.Data = om // or false

	return c.JSON(http.StatusOK, _res)
}

// otp action form
func (r *RegisterHandler) Otp(c echo.Context) error {

	fmt.Println("Otp")
	otpModel := models.OtpModel{}
	c.Bind(&otpModel)
	// _otp, _ := strconv.Atoi(c.FormValue("otp"))
	// otpModel.Otp = _otp                           // get params
	// otpModel.Recipient = c.FormValue("recipient") // get params
	// otpModel.Type = c.FormValue("type")           // get params

	id, err := r.Repo.Otp(otpModel)

	if err != nil {
		fmt.Println("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}
	fmt.Println("come1", id, *id)
	if id != nil && *id == 0 {
		fmt.Println("come")
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		msg := map[string]string{
			"otp": "Otp does not match.",
		}
		_res.ErrorMessage = msg
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "500"
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	otp_res := models.OtpModel{Success: true, Otp: otpModel.Otp}
	_res.Data = otp_res // or false

	return c.JSON(http.StatusOK, _res)
}
