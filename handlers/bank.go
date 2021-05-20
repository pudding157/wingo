package handlers

import (
	"fmt"
	"winapp/models"

	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func BankHandler(db *gorm.DB) *Handler {

	Bank := []models.Bank{
		{Id: 1, Name: "ธนาคารกสิกรไทย จำกัด (มหาชน)", Is_active: true},
		{Id: 2, Name: "ธนาคารไทยพาณิชย์ จำกัด (มหาชน)", Is_active: true},
		{Id: 3, Name: "ธนาคารกรุงเทพ จำกัด (มหาชน)", Is_active: true},
		{Id: 4, Name: "ธนาคารกรุงศรีอยุธยา จำกัด (มหาชน)", Is_active: true},
		{Id: 5, Name: "ธนาคารกรุงไทย จำกัด (มหาชน)", Is_active: true},
		{Id: 6, Name: "ธนาคารนครหลวงไทย จำกัด (มหาชน)", Is_active: true},
		{Id: 7, Name: "ธนาคารทหารไทย จำกัด (มหาชน)", Is_active: true},
		{Id: 8, Name: "ธนาคารยูโอบี จำกัด (มหาชน)", Is_active: true},
		{Id: 9, Name: "ธนาคารออมสิน", Is_active: true},
		{Id: 10, Name: "ธนาคารอาคารสงเคราะห์", Is_active: true},
		{Id: 11, Name: "ธนาคารซีไอเอ็มบีไทย จำกัด (มหาชน)", Is_active: true},
		{Id: 12, Name: "ธนาคารธนชาต จำกัด (มหาชน)", Is_active: true},
		{Id: 13, Name: "ธนาคารเพื่อการเกษตรและสหกรณ์การเกษตร (มหาชน)", Is_active: true},
	}
	if !db.HasTable(Bank) {
		fmt.Println("No table")
		db.AutoMigrate(&Bank) // สร้าง table, field ต่างๆที่ไม่เคยมี
		for _, _bank := range Bank {
			db.Create(_bank)
		}
		fmt.Println("migrate data bank and create bank")
	} else {
		fmt.Println("has table")
	}
	return &Handler{DB: db}
}

type BankModel struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// otp/send action form
func (h *Handler) Get_all_bank(c echo.Context) error {

	fmt.Println("Get all bank")

	banks := []models.Bank{}
	bankModel := []BankModel{}

	if err := h.DB.Find(&banks).Error; err != nil {
		fmt.Println("h.DB.Find(&banks) => ", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	for _, _bank := range banks {

		_bankModel := BankModel{Id: _bank.Id, Name: _bank.Name}
		bankModel = append(bankModel, _bankModel)
	}
	_res := models.Response{}
	_res.Data = bankModel // or false

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return c.JSON(http.StatusOK, _res)
}
