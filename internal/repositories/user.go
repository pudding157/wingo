package repositories

import (
	"fmt"
	"winapp/internal/app"
	"winapp/internal/models"
	"winapp/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

type UserRepository interface {
	GetProfile(token string) (*models.UserProfile, error)
}

type UserRepo struct {
	c *app.Config
}

func NewUserRepo(c *app.Config) *UserRepo {
	return &UserRepo{c: c}
}

func (r *UserRepo) GetProfile(token string) (*models.UserProfile, error) {

	// auth_header := c.Request().Header.Get("Authorization")
	// auth_len := len(auth_header)
	// token := auth_header[7:auth_len]

	fmt.Println("token :", token)

	claims := jwt.MapClaims{}

	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || t == nil {
		fmt.Println("token err", err)
		return nil, err
	}
	// do something with decoded claims
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }
	userid := claims["user_id"]
	// userid := c.Param("userid")
	fmt.Println("userid :", userid)
	User := models.User{}
	r.c.DB.Where("id = ?", userid).Find(&User)
	User_bank := models.User_bank{}
	r.c.DB.Where("user_id = ?", User.Id).Find(&User_bank)

	Bank := models.Bank{}
	r.c.DB.Where("id = ?", User_bank.Bank_id).Find(&Bank)

	User_Profile := models.UserProfile{}
	User_Profile.Name = User.First_name + " " + User.Last_name
	User_Profile.PhoneNumber = User.Phone_number
	User_Profile.BankAccount = User_bank.Bank_account
	User_Profile.BankName = Bank.Name
	User_Profile.Status = utils.MEMBER.String()

	return &User_Profile, nil
}
