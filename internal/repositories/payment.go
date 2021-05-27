package repositories

import (
	// "errors"
	// "fmt"
	// "log"
	// "time"
	"winapp/internal/app"
	// "winapp/internal/models"
	// "winapp/internal/utils"
	// "github.com/dgrijalva/jwt-go"
)

type PaymentRepository interface {
	Deposit() error
	Withdraw() error
}

type PaymentRepo struct {
	c *app.Config
}

func NewPaymentRepo(c *app.Config) *PaymentRepo {
	return &PaymentRepo{c: c}
}

func (r *PaymentRepo) Deposit() error {
	return nil
}

func (r *PaymentRepo) Withdraw() error {
	return nil
}
