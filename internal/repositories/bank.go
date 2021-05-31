package repositories

import (
	"fmt"
	"winapp/internal/app"
	"winapp/internal/models"
)

type BankRepository interface {
	GetBanks() ([]models.Bank, error)
	GetAdminBanks() ([]models.Admin_Bank, error)
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

	// get only isactive
	if err := r.c.DB.Find(&b, "is_active = 1").Error; err != nil {
		fmt.Println("h.DB.Find(&banks) => ", err)
		return nil, err
	}
	fmt.Println("h.DB.Find(&banks) => true =>  ", b)
	return b, nil
}

func (r *BankRepo) GetAdminBanks() ([]models.Admin_Bank, error) {

	fmt.Println("Get all admin bank")

	b := []models.Admin_Bank{}

	if err := r.c.DB.Find(&b).Error; err != nil {
		fmt.Println("h.DB.Find(&banks) => ", err)
		return nil, err
	}
	banks, err := r.GetBanks()
	if err != nil {
		return nil, err
	}

	for i, _b := range b {
		bn := ""
		for _, bb := range banks {

			if bb.Id == _b.BankId {
				bn = bb.Name
				break
			}
		}
		b[i].BankName = bn
	}
	fmt.Println("b => ", b)
	return b, nil
}
