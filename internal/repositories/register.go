package repositories

import (
	"errors"
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
	c  *app.Config
	lr *LoginRepo
}

func NewRegisterRepo(c *app.Config, lr *LoginRepo) *RegisterRepo {
	return &RegisterRepo{c: c, lr: lr}
}

func (r *RegisterRepo) Register(Bind_registerFormModel models.RegisterFormModel) (*string, error) {

	fmt.Println("Bind_registerFormModel, ", Bind_registerFormModel)

	if len(Bind_registerFormModel.Password) < 8 {
		return nil, errors.New("Password lower than 8 characters.")
	}
	if Bind_registerFormModel.Bank_id == 0 {
		return nil, errors.New("Please select the bank.")
	}

	User_bank := models.User_Bank{}
	err := r.c.DB.Where("bank_account = ?", Bind_registerFormModel.Bank_account).Find(&User_bank).Error
	if err == nil {
		fmt.Println("bank account already exist", User_bank)
		return nil, errors.New("bank account already exist.")
	}

	_passwordHashed := utils.HashStr(Bind_registerFormModel.Password)

	fmt.Println("Hash is => ", _passwordHashed)

	u := models.User{}

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	s := utils.StringWithCharset(16, charset)

	u.Affiliate = s
	u.First_name = Bind_registerFormModel.First_name
	u.Last_name = Bind_registerFormModel.Last_name
	u.Phone_number = Bind_registerFormModel.Phone_number
	u.Username = Bind_registerFormModel.Username
	u.Password = _passwordHashed
	_now := time.Now().UTC().Format(time.RFC3339)
	u.Created_at = _now
	u.Updated_at = _now
	u.Registration_otp = strconv.Itoa(Bind_registerFormModel.Otp)

	if err := r.c.DB.Save(&u).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
	}
	User_bank.BankId = Bind_registerFormModel.Bank_id
	User_bank.BankAccount = Bind_registerFormModel.Bank_account
	User_bank.UserId = u.Id
	User_bank.Created_at = _now

	if err := r.c.DB.Save(&User_bank).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
	}

	// gen token for start login
	t, err := r.lr.GenToken(u)
	if err != nil {
		return nil, err
	}

	return t, nil
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

	history := models.Otp_History{}
	history.Type = utils.PHONE_NUMBER.Index()
	history.SendTo = phone_number
	history.Otp = otp
	history.CreatedAt = time.Now().UTC().Format(time.RFC3339)
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

	fmt.Println("Otp", otpModel)

	keyType, _err := utils.EnumFromKey(otpModel.Type, utils.GetEnumArray("otp"))
	if _err != nil {
		log.Fatal(_err)
	}
	history := models.Otp_History{}
	err := r.c.DB.Where("otp = ? AND type = ? AND send_to = ?", otpModel.Otp, keyType.Index(), otpModel.Recipient).Find(&history).Error

	if history.Id == 0 || err != nil {
		return &history.Id, err
	}

	return &history.Id, nil
}
