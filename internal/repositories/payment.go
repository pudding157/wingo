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
	Transactions(t models.LoadMoreModel) ([]models.User_Bind_History, int64, error)
}

type PaymentRepo struct {
	c *app.Config
}

func NewPaymentRepo(c *app.Config) *PaymentRepo {
	return &PaymentRepo{c: c}
}

func (r *PaymentRepo) Deposit(uh models.User_History) error {
	// keyType, _err := utils.EnumFromIndex(uh.Status, utils.GetEnumArray("depositStatus"))
	_now := time.Now().UTC() //.Format(time.RFC3339)

	fmt.Println("_now => ", _now)
	uh.UserId = r.c.UI
	uh.Type = utils.DEPOSIT.Index()
	uh.CreatedAt = _now
	uh.UpdatedAt = _now
	fmt.Println("uh => ", uh)
	if err := r.c.DB.Save(&uh).Error; err != nil {
		log.Print("err => ", err)
		return err
	}
	return nil

}

func (r *PaymentRepo) Withdraw(uh models.User_History) error {
	fmt.Println("amount => ", uh.Amount)
	_now := time.Now().UTC() //.Format(time.RFC3339)

	if uh.Amount < 500 {
		return errors.New("Amount is lower than the withdrawal amount.")
	}

	uh.TransferredAt = _now
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

func (r *PaymentRepo) Transactions(t models.LoadMoreModel) ([]models.User_Bind_History, int64, error) {
	fmt.Println("transaction type => ", t)
	kt, err := utils.EnumFromKey(t.Type, utils.GetEnumArray("transactionType"))
	if err != nil {
		return nil, int64(0), errors.New("no transactions type")
	}
	uh := []models.User_History{}
	cs := "user_id = " + strconv.Itoa(r.c.UI)
	if t.Type != "all" {
		cs += " and type = " + strconv.Itoa(kt.Index())
	}

	u := []models.User_Bind_History{}
	c := 0
	if err := r.c.DB.Model(&uh).Where(cs).Count(&c).Limit(t.Take).Offset(t.Skip).Order("created_at desc").Find(&uh, cs).Error; err != nil {
		fmt.Println("h.DB.Find User_History error => ", err)
		return nil, int64(0), err
	}
	for _, uhs := range uh {
		ubh := models.User_Bind_History{}
		ubh.AdminBankAccount = uhs.AdminBankAccount
		ubh.Amount = uhs.Amount
		ubh.Type = uhs.Type
		ubh.TransferredAt = uhs.TransferredAt.String()
		ubh.Status = uhs.Status

		u = append(u, ubh)
		// b[i].BankName = bn
	}

	return u, int64(c), nil
}
