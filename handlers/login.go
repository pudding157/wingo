package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"winapp/models"
	"winapp/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func LoginHandler(db *gorm.DB) *Handler {

	User_login := models.User_login{}
	if !db.HasTable(User_login) {
		fmt.Println("No table")
		db.AutoMigrate(&User_login) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data User_bank")
	}
	return &Handler{DB: db}
}

type jwtCustomClaims struct {
	User_id string `json:"user_id"`
	jwt.StandardClaims
}

type redisValue struct {
	User_id     int    `json:"user_id"`
	Expire_date string `json:"expire_date"`
}

func (h *Handler) Login(c echo.Context) error {

	Bind_user := &models.User{}
	c.Bind(&Bind_user)
	fmt.Println("Bind_user, ", Bind_user)

	User := models.User{}
	h.DB.Where("username = ?", Bind_user.Username).Find(&User)
	fmt.Println("user => ", User)
	if !utils.DehashStr(User.Password, Bind_user.Password) {
		return echo.ErrUnauthorized
	}

	// claims := &jwtCustomClaims{
	// 	User_id: strconv.Itoa(User.Id),
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add((time.Hour * 8760) * 2).Unix(),
	// 	},
	// }
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = strconv.Itoa(User.Id)
	claims["exp"] = time.Now().Add((time.Hour * 8760) * 2).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	// create token and will add to redis too
	User_login := &models.User_login{}
	User_login.User_id = User.Id
	User_login.Username = User.Username
	_now := time.Now().Format(time.RFC3339)
	User_login.Created_at = _now
	User_login.Token = t
	if err := h.DB.Save(&User_login).Error; err != nil {
		fmt.Println("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "400"
		return c.JSON(http.StatusBadRequest, _res)
	}
	fmt.Println("1111111111111")
	r := RedisHandler{}
	r.Initialize()
	rt := time.Unix(time.Now().Add((time.Hour*8760)*2).Unix(), 0)

	redisValue := redisValue{}
	redisValue.User_id = User.Id
	redisValue.Expire_date = time.Now().Add(time.Second).Format(time.RFC3339)
	rv, _ := json.Marshal(redisValue)
	errAccess := r.DB.Set(t, string(rv), rt.Sub(time.Now())).Err()
	fmt.Println("222222222222222222", redisValue)
	if errAccess != nil {
		fmt.Println("2333333333333333333 ", errAccess)
		return errAccess
	}
	// 	type User_login struct {
	// 	Token       string `gorm:"primary_key" json:"token"`
	// 	Id          int    `gorm:"type:autoIncrement" json:"id"`
	// 	User_id     int    `json:"user_id"`
	// 	User        User   `gorm:"foreignKey:User_id"`
	// 	Username    string `gorm:"not null" json:"username"`
	// 	Ip_address  string `gorm:"not null" json:"ip_address"`
	// 	Mac_address string `gorm:"not null" json:"mac_address"`
	// 	User_agent  string `gorm:"not null" json:"user_agent"`
	// 	Created_at  string `json:"created_at"`
	// }

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func (h *Handler) restricted(c echo.Context) error {
	userid := c.Param("userid")
	fmt.Println("userid :", userid)
	User := models.User{}
	h.DB.Where("id = ?", userid).Find(&User)
	j, _ := json.Marshal(User)
	return c.String(http.StatusOK, string(j))
}
