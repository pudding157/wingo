package repositories

import (
	"errors"
	"fmt"
	"log"
	"time"
	"winapp/internal/app"
	"winapp/internal/models"
	"winapp/internal/utils"
	// "github.com/dgrijalva/jwt-go"
)

type UserRepository interface {
	GetProfile() (*models.UserProfile, error)
	ChangePassword(ph models.Password_History) (*string, error)
	GetAffiliate() (*string, error)
}

type UserRepo struct {
	c  *app.Config
	lr *LoginRepo
}

func NewUserRepo(c *app.Config, lr *LoginRepo) *UserRepo {
	return &UserRepo{c: c, lr: lr}
}

func (r *UserRepo) GetProfile() (*models.UserProfile, error) {
	User := models.User{}
	fmt.Println("get profile", r.c)

	err := r.c.DB.Where("id = ?", r.c.UI).Find(&User).Error
	if err != nil {
		fmt.Println("User DB => ", err)
		return nil, err
	}

	User_bank := models.User_Bank{}

	err = r.c.DB.Where("user_id = ?", User.Id).Find(&User_bank).Error
	if err != nil {
		fmt.Println("User Bank DB => ", err)
		return nil, err
	}

	Bank := models.Bank{}
	err = r.c.DB.Where("id = ?", User_bank.BankId).Find(&Bank).Error
	if err != nil {
		fmt.Println("Bank DB => ", err)
		return nil, err
	}

	User_Profile := models.UserProfile{}
	User_Profile.Username = User.Username
	User_Profile.Name = User.FirstName + " " + User.LastName
	User_Profile.PhoneNumber = User.PhoneNumber
	User_Profile.BankAccount = User_bank.BankAccount
	User_Profile.BankName = Bank.Name
	es := utils.GetEnumArray("userStatus")
	fmt.Println(es)
	mt, _err := utils.EnumFromIndex(User.Status, es)
	if _err != nil {
		fmt.Println("EnumFromIndex(User.Status, es) => ", _err)
	}

	User_Profile.Status = mt.String(es)

	cu := []models.User{}
	err = r.c.DB.Where("parent_user_id = ?", User.Id).Find(&cu).Error
	if err != nil {
		fmt.Println("User DB parent_user_id => ", err)
		return nil, err
	} else {
		for _, _u := range cu {
			name := utils.HiddenLastString(4, _u.Username)
			User_Profile.ChildUserNames = append(User_Profile.ChildUserNames, name)
		}
	}
	fmt.Println("before User.ParentUserId", User.ParentUserId)
	if User.ParentUserId != nil && *User.ParentUserId != 0 {
		fmt.Println("after User.ParentUserId", User.ParentUserId)
		pu := models.User{}
		err := r.c.DB.Where("id = ?", User.ParentUserId).Find(&pu).Error
		if err != nil {
			fmt.Println("User DB => ", err)
			return nil, err
		}
		re := utils.HiddenLastString(4, pu.Username)
		// re := regexp.MustCompile(`\w{4}$`).ReplaceAllString(pu.Username, "")
		User_Profile.ParentUserName = re
	}

	fmt.Println("end User_Profile => ", User_Profile)

	return &User_Profile, nil
}

func (r *UserRepo) ChangePassword(ph models.Password_History) (*string, error) {

	//Password_History

	u := models.User{}
	err := r.c.DB.Where("id = ?", r.c.UI).Find(&u).Error
	fmt.Println("user => ", u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if !utils.DehashStr(u.Password, ph.OldPassword) {
		return nil, errors.New("Password does not match.")
	}

	ph.Username = u.Username
	ph.IPAddress = ""
	ph.MACAddress = ""
	ph.Browser = ""
	_now := time.Now().UTC() //.Format(time.RFC3339)
	ph.CreatedAt = _now
	ph.OldPassword = u.Password

	np := ph.NewPassword

	_pwd := utils.HashStr(np)
	fmt.Println("old => ", u.Password)
	u.Password = _pwd

	if err := r.c.DB.Save(&u).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
	}

	fmt.Println("newest => ", u.Password)
	ph.NewPassword = _pwd

	if err := r.c.DB.Save(&ph).Error; err != nil {
		log.Print("err => ", err)
		return nil, err
	}

	err = r.lr.Logout() // force logout and clear token
	if err != nil {
		return nil, err
	}

	t, err := r.lr.GenToken(u)
	if err != nil {
		return nil, err
	}
	return t, nil
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (r *UserRepo) GetAffiliate() (*string, error) {

	u := models.User{}
	err := r.c.DB.Where("id = ?", r.c.UI).Find(&u).Error
	fmt.Println("user => ", u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if u.Affiliate == "" {
		s := utils.StringWithCharset(16, charset)
		r.c.DB.Model(&u).Updates(models.User{UpdatedAt: time.Now().UTC(), Affiliate: s})
		return &s, nil
	}

	return &u.Affiliate, nil
}
