package handlers

import (
	"fmt"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type BankHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.BankRepository
}

func NewBankHandler(repo repositories.BankRepository) *BankHandler {
	return &BankHandler{Repo: repo}
}

// otp/send action form
func (r *BankHandler) GetBanks(c echo.Context) error {

	fmt.Println("Get all bank")

	banks, err := r.Repo.GetBanks()

	if err != nil {
		// 	fmt.Println("h.DB.Find(&banks) => ", err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	// bankModel := []BankModel{}

	// if err := h.c.DB.Find(&banks).Error; err != nil {
	// 	fmt.Println("h.DB.Find(&banks) => ", err)
	// 	return c.JSON(http.StatusBadRequest, err)
	// }
	// for _, _bank := range banks {

	// 	_bankModel := BankModel{Id: _bank.Id, Name: _bank.Name}
	// 	bankModel = append(bankModel, _bankModel)
	// }
	_res := models.Response{}
	_res.Data = banks // or false

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.JSON(http.StatusOK, _res)
}
