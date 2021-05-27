package app

import (
	"fmt"
	"winapp/internal/models"

	"github.com/jinzhu/gorm"
	// "github.com/spf13/viper"
)

type Config_db struct {
	DB *gorm.DB
}

func (c *Config_db) Init() {
	c.migrate_bank()
	c.migrate_User()
	c.migrate_other()
}

func (c *Config_db) migrate_bank() {

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
	if !c.DB.HasTable(Bank) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Bank) // สร้าง table, field ต่างๆที่ไม่เคยมี
		for _, _bank := range Bank {
			c.DB.Create(_bank)
		}
		fmt.Println("migrate data bank and create bank")
	} else {
		fmt.Println("has table")
	}
}

func (c *Config_db) migrate_User() {

	User_login := models.User_login{}
	if !c.DB.HasTable(User_login) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&User_login) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data User_login")
	}

	User := []models.User{}
	if !c.DB.HasTable(User) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&User) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data User")
	} else {
		// c.DB.AutoMigrate(&User) // สร้าง table, field ต่างๆที่ไม่เคยมี
	}

	User_bank := models.User_bank{}
	if !c.DB.HasTable(User_bank) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&User_bank) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data User_bank")
	} else {
		// ถ้าจะเพิ่ม unique ถ้า ในตารางมีข้อมูลซ้ำจะไม่สามารถทำได้
		c.DB.Model(&User_bank).AddUniqueIndex("bank_account", "bank_account")
	}

	ph := models.Password_History{}
	if !c.DB.HasTable(ph) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&ph) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data password history")
	} else {
		// c.DB.AutoMigrate(&User) // สร้าง table, field ต่างๆที่ไม่เคยมี
	}
}

func (c *Config_db) migrate_other() {

	Otp_history := []models.Otp_history{}
	if !c.DB.HasTable(Otp_history) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Otp_history) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data Otp_history")
	}
}
