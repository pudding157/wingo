package handlers

import (
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

	return &Handler{DB: db}
}

const secret = "secret"

type jwtCustomClaims struct {
	User_id string `json:"user_id"`
	jwt.StandardClaims
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

	claims := &jwtCustomClaims{
		User_id: strconv.Itoa(User.Id),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add((time.Hour * 8760) * 2).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("err", err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func (h *Handler) restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	fmt.Println(claims, "claims")
	user_id := claims.User_id
	return c.String(http.StatusOK, "Welcome "+user_id+"!")
}
