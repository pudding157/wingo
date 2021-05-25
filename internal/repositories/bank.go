package repositories

import (
	"fmt"
	"winapp/internal/app"
	"winapp/internal/models"
)

type BankRepository interface {
	GetBanks() ([]models.Bank, error)
}

type BankRepo struct {
	c *app.Config
}

func NewBankRepo(c *app.Config) *BankRepo {
	return &BankRepo{c: c}
}

func (r *BankRepo) GetBanks() ([]models.Bank, error) {

	fmt.Println("Get all bank")

	b := []models.Bank{}

	if err := r.c.DB.Find(&b).Error; err != nil {
		fmt.Println("h.DB.Find(&banks) => ", err)
		return nil, err
	}

	return b, nil
}
