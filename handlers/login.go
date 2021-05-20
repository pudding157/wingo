package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func LoginHandler(db *gorm.DB) *Handler {

	return &Handler{DB: db}
}

const secret = "secret"

type jwtCustomClaims struct {
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func (h *Handler) Login(c echo.Context) error {

	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "pieter" || password != "claerhout" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		Name:  "Pieter Claerhout",
		UUID:  "9E98C454-C7AC-4330-B2EF-983765E00547",
		Admin: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	fmt.Println(claims, "claims")
	name := claims.Name
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
