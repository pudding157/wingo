package repositories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"winapp/internal/app"
	"winapp/internal/models"
	"winapp/internal/utils"
)

type PaymentRepository interface {
	Deposit(uh models.User_History) error
	Withdraw(uh models.User_History) error
	Transactions(t string) ([]models.User_History, error)
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

func (r *PaymentRepo) Withdraw(uh models.User_History) error {
	fmt.Println("amount => ", uh.Amount)
	_now := time.Now().Format(time.RFC3339)

	if uh.Amount < 500 {
		return errors.New("Amount is lower than the withdrawal amount.")
	}
	uh.UserId = r.c.UI
	uh.Type = utils.WITHDRAW.Index()
	uh.CreatedAt = _now
	uh.UpdatedAt = _now
	uh.Status = utils.AWAITING.Index()
	fmt.Println("uh => ", uh)
	if err := r.c.DB.Save(&uh).Error; err != nil {
		log.Print("err => ", err)
		// return nil, err
	}

	return nil
}

func (r *PaymentRepo) Transactions(t string) ([]models.User_History, error) {
	fmt.Println("transaction type => ", t)
	kt, _err := utils.EnumFromKey(t, utils.GetEnumArray("transactionType"))
	if _err != nil {
		return nil, errors.New("no transactions type")
	}
	uh := []models.User_History{}
	cs := "user_id = " + strconv.Itoa(r.c.UI)
	if t != "all" {
		cs += " and type = " + strconv.Itoa(kt.Index())
	}
	if err := r.c.DB.Find(&uh, cs).Error; err != nil {
		fmt.Println("h.DB.Find User_History => ", err)
		return nil, err
	}

	return uh, nil
}
