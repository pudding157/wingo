package utils

import (
	"math/rand"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"winapp/models"
)

// JSONResponse func
func JSONResponse(c echo.Context, data interface{}, err error) error {
	code, message := "200", "OK"
	if err != nil {
		code, message = "500", err.Error()
		if err.Error() == "409" {
			code, message = "409", "Some data already exist."
		} else if err.Error() == "404" {
			code, message = "404", "Data not found."
		} else if err.Error() == "401" {
			code, message = "401", "User Unauthorized."
		}
	}
	res := models.Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return c.JSON(http.StatusOK, res)
}

// GeneratePassword func
func GeneratePassword() string {
	var (
		lowerCharSet   = "abcdedfghijklmnopqrst"
		upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*"
		numberSet      = "0123456789"
		allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
		minSpecialChar = 1
		minNum         = 1
		minUpperCase   = 1
		passwordLength = 12
	)
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
