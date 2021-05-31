package repositories

import (
	"fmt"
	"log"
	"time"
	"winapp/internal/app"
	"winapp/internal/models"
	"winapp/internal/utils"
)

type PaymentRepository interface {
	Deposit(uh models.User_History) error
	Withdraw() error
	Transactions(t string) error
}

type PaymentRepo struct {
	c *app.Config
}

func NewPaymentRepo(c *app.Config) *PaymentRepo {
	return &PaymentRepo{c: c}
}

func (r *PaymentRepo) Deposit(uh models.User_History) error {
	// keyType, _err := utils.EnumFromIndex(uh.Status, utils.GetEnumArray("depositStatus"))
	_now := time.Now().Format(time.RFC3339)

	uh.UserId = r.c.UI
	uh.Type = utils.DEPOSIT.Index()
	uh.CreatedAt = _now
	uh.UpdatedAt = _now
	fmt.Println("uh => ", uh)
	if err := r.c.DB.Save(&uh).Error; err != nil {
		log.Print("err => ", err)
		// return nil, err
	}
	return nil

}

func (r *PaymentRepo) Withdraw() error {
	return nil
}

func (r *PaymentRepo) Transactions(t string) error {
	fmt.Println("transaction type => ", t)
	return nil
}
