package repositories

import (
	"fmt"
	"log"
	"strconv"
	"time"
	"winapp/internal/app"
	"winapp/internal/models"
	"winapp/internal/utils"
)

type RegisterRepository interface {
	Register(Bind_registerFormModel models.RegisterFormModel) (*string, error)
	Otp_send(phone_number string) (*models.OtpModel, error)
	Otp(otpModel models.OtpModel) (*int, error)
}

type RegisterRepo struct {
	c *app.Config
}

func NewRegisterRepo(c *app.Config) *RegisterRepo {
	return &RegisterRepo{c: c}
}

func (r *RegisterRepo) Register(Bind_registerFormModel models.RegisterFormModel) (*string, error) {

	fmt.Println("Bind_registerFormModel, ", Bind_registerFormModel)

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

	if err := r.c.DB.Save(&User).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
	}

	User_bank := models.User_bank{}
	User_bank.Bank_id = Bind_registerFormModel.Bank_id
	User_bank.Bank_account = Bind_registerFormModel.Bank_account
	User_bank.User_id = User.Id
	User_bank.Created_at = _now

	if err := r.c.DB.Save(&User_bank).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
	}
	token := "asd"

	return &token, nil
}

func (r *RegisterRepo) Otp_send(phone_number string) (*models.OtpModel, error) {

	fmt.Println("Otp send")

	// phone_number := c.FormValue("phone_number") // get params
	// fmt.Println("phone_number => ", phone_number)

	// todo otp
	otp := 123456

	_res := models.Response{}
	otp_res := models.OtpModel{}
	otp_res.Success = true
	otp_res.Otp = otp // get params
	_res.Data = otp_res

	history := models.Otp_history{}
	history.Type = utils.PHONE_NUMBER.Index()
	history.Send_to = phone_number
	history.Otp = otp
	history.Created_at = time.Now().Format(time.RFC3339)
	fmt.Println("history => ", history)

	if err := r.c.DB.Save(&history).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
		// _res := models.ErrorResponse{}
		// _res.Error = "Validation Failed"
		// // _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		// _res.Error_code = "400"
		// return c.JSON(http.StatusBadRequest, _res)
	}
	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &otp_res, nil
}

func (r *RegisterRepo) Otp(otpModel models.OtpModel) (*int, error) {

	fmt.Println("Otp")

	keyType, _err := utils.EnumFromKey(otpModel.Type)
	if _err != nil {
		log.Fatal(_err)
	}
	history := models.Otp_history{}
	err := r.c.DB.Where("otp = ? AND type = ? AND send_to = ?", otpModel.Otp, keyType.Index(), otpModel.Recipient).Find(&history).Error

	if history.Id == 0 || err != nil {
		return &history.Id, err
	}

	return &history.Id, nil
}
