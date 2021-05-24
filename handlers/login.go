package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"winapp/app"
	"winapp/models"
	"winapp/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	// DB *gorm.DB
	// R  *redis.Client
	c *app.Config
}

// func LoginHandler(c *app.Config) *Handler {

// 	User_login := models.User_login{}
// 	if !c.DB.HasTable(User_login) {
// 		fmt.Println("No table")
// 		c.DB.AutoMigrate(&User_login) // สร้าง table, field ต่างๆที่ไม่เคยมี
// 		fmt.Println("migrate data User_login")
// 	}
// 	return &Handler{DB: c.DB, R: c.R}
// }

type jwtCustomClaims struct {
	User_id string `json:"user_id"`
	jwt.StandardClaims
}

type redisValue struct {
	User_id     int    `json:"user_id"`
	Expire_date string `json:"expire_date"`
}

func (h *LoginHandler) Login(c echo.Context) error {

	Bind_user := &models.User{}
	c.Bind(&Bind_user)
	fmt.Println("Bind_user, ", Bind_user)

	User := models.User{}
	h.c.DB.Where("username = ?", Bind_user.Username).Find(&User)
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
	if err := h.c.DB.Save(&User_login).Error; err != nil {
		fmt.Println("err => ", err)
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = "400"
		return c.JSON(http.StatusBadRequest, _res)
	}
	fmt.Println("1111111111111")
	r := h.c.R
	rt := time.Unix(time.Now().Add((time.Hour*8760)*2).Unix(), 0)

	redisValue := redisValue{}
	redisValue.User_id = User.Id
	redisValue.Expire_date = time.Now().Add(time.Hour * 2).Format(time.RFC3339)
	rv, _ := json.Marshal(redisValue)
	errAccess := r.Set(t, string(rv), rt.Sub(time.Now())).Err()
	fmt.Println("222222222222222222", redisValue)
	if errAccess != nil {
		fmt.Println("2333333333333333333 ", errAccess)
		return errAccess
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
