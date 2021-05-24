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
	// bm := []BankModel{}

	if err := r.c.DB.Find(&b).Error; err != nil {
		fmt.Println("h.DB.Find(&banks) => ", err)
		return nil, err
	}

	return b, nil
	// // for _, _bank := range b {

	// // 	_bankModel := BankModel{Id: _bank.Id, Name: _bank.Name}
	// // 	bm = append(bm, _bankModel)
	// // }
	// _res := models.Response{}
	// _res.Data = b // or false

	// // if err != nil {
	// // 	log.Fatal(err)
	// // }

	// return c.JSON(http.StatusOK, _res)
}
