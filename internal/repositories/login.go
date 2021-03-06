package repositories

import (
	// "fmt"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"time"
	"winapp/internal/app"
	"winapp/internal/models"
	"winapp/internal/utils"

	"github.com/dgrijalva/jwt-go"
)

type LoginRepository interface {
	Login(bu models.User) (*string, error)
	Logout() error
}

type LoginRepo struct {
	c *app.Config
}

func NewLoginRepo(c *app.Config) *LoginRepo {
	return &LoginRepo{c: c}
}

func (r *LoginRepo) Login(bu models.User) (*string, error) {

	fmt.Println("Bind_user, ", bu)

	u := models.User{}
	r.c.DB.Where("username = ?", bu.Username).Find(&u)
	fmt.Println("user => ", u)
	if !utils.DehashStr(u.Password, bu.Password) {
		// return nil, echo.ErrUnauthorized
		return nil, errors.New("Password is incorrect.")
	}

	t, err := r.GenToken(u)
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	return t, nil
}

func (r *LoginRepo) GenToken(u models.User) (*string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	cl := token.Claims.(jwt.MapClaims)
	cl["user_id"] = strconv.Itoa(u.Id)
	cl["exp"] = time.Now().UTC().Add((time.Hour * 8760) * 2).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	// create token and will add to redis too
	ul := &models.User_Login{}
	ul.UserId = u.Id
	ul.Username = u.Username
	_now := time.Now().UTC()
	ul.CreatedAt = _now
	ul.Token = t
	if err := r.c.DB.Save(&ul).Error; err != nil {
		fmt.Println("err => ", err)
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		// _res.Error_code = "400"
		return nil, err
	}
	rt := time.Unix(time.Now().UTC().Add((time.Hour*8760)*2).Unix(), 0)

	rvm := models.RedisValue{}
	rvm.UserId = u.Id

	es := utils.GetEnumArray("userStatus")
	mt, _err := utils.EnumFromIndex(u.Status, es)
	if _err != nil {
		fmt.Println("EnumFromIndex(u.Status, es) => ", _err)
	}
	rvm.RoleName = mt.String(es)

	rvm.ExpireDate = time.Now().UTC().Add((time.Hour * 8760) * 2).Format(time.RFC3339)
	rv, _ := json.Marshal(rvm)
	err = r.c.R.Set(t, string(rv), rt.Sub(time.Now().UTC())).Err()
	fmt.Println("222222222222222222", rvm)
	if err != nil {
		fmt.Println("2333333333333333333 ", err)
		return nil, err
	}
	return &t, nil
}

func (r *LoginRepo) Logout() error {
	fmt.Println("token => ", r.c.T)

	err := r.c.R.Del(r.c.T).Err()
	if err != nil {
		fmt.Println("token err", err)
		return err
	}

	return nil
}
