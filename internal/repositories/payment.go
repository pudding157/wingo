package repositories

import (
	// "errors"
	// "fmt"
	// "log"
	// "time"
	"fmt"
	"winapp/internal/app"
	// "winapp/internal/models"
	// "winapp/internal/utils"
	// "github.com/dgrijalva/jwt-go"
)

type PaymentRepository interface {
	Deposit() error
	Withdraw() error
	Transactions(t string) error
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

func (r *PaymentRepo) Transactions(t string) error {
	fmt.Println("transaction type => ", t)
	return nil
}
