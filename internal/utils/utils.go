package utils

import (
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
	"winapp/internal/models"

	"github.com/labstack/echo/v4"
)

// JSONResponse func
func JSONResponse(c echo.Context, data interface{}, err error) error {
	code, message := strconv.Itoa(http.StatusOK), "OK"
	if err != nil {
		code, message = strconv.Itoa(http.StatusInternalServerError), err.Error()
		if err.Error() == strconv.Itoa(http.StatusConflict) {
			code, message = strconv.Itoa(http.StatusConflict), "Some data already exist."
		} else if err.Error() == strconv.Itoa(http.StatusNotFound) {
			code, message = strconv.Itoa(http.StatusNotFound), "Data not found."
		} else if err.Error() == strconv.Itoa(http.StatusUnauthorized) {
			code, message = strconv.Itoa(http.StatusUnauthorized), "User Unauthorized."
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

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	seededRand := rand.New(
		rand.NewSource(time.Now().UTC().UnixNano()))
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
