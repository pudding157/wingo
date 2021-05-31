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
		{Id: 1, Name: "ธนาคารกสิกรไทย จำกัด (มหาชน)", IsActive: true},
		{Id: 2, Name: "ธนาคารไทยพาณิชย์ จำกัด (มหาชน)", IsActive: true},
		{Id: 3, Name: "ธนาคารกรุงเทพ จำกัด (มหาชน)", IsActive: true},
		{Id: 4, Name: "ธนาคารกรุงศรีอยุธยา จำกัด (มหาชน)", IsActive: true},
		{Id: 5, Name: "ธนาคารกรุงไทย จำกัด (มหาชน)", IsActive: true},
		{Id: 6, Name: "ธนาคารนครหลวงไทย จำกัด (มหาชน)", IsActive: true},
		{Id: 7, Name: "ธนาคารทหารไทย จำกัด (มหาชน)", IsActive: true},
		{Id: 8, Name: "ธนาคารยูโอบี จำกัด (มหาชน)", IsActive: true},
		{Id: 9, Name: "ธนาคารออมสิน", IsActive: true},
		{Id: 10, Name: "ธนาคารอาคารสงเคราะห์", IsActive: true},
		{Id: 11, Name: "ธนาคารซีไอเอ็มบีไทย จำกัด (มหาชน)", IsActive: true},
		{Id: 12, Name: "ธนาคารธนชาต จำกัด (มหาชน)", IsActive: true},
		{Id: 13, Name: "ธนาคารเพื่อการเกษตรและสหกรณ์การเกษตร (มหาชน)", IsActive: true},
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

	Admin_Bank := []models.Admin_Bank{
		{Id: 1, BankId: 4, BankAccount: "6382625487", IsActive: true},
		{Id: 2, BankId: 3, BankAccount: "4552113322", IsActive: true},
		{Id: 3, BankId: 5, BankAccount: "9776665544", IsActive: true},
	}
	if !c.DB.HasTable(Admin_Bank) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Admin_Bank) // สร้าง table, field ต่างๆที่ไม่เคยมี
		for _, _bank := range Admin_Bank {
			c.DB.Create(_bank)
		}
		fmt.Println("migrate data bank and create Admin_Bank")
	} else {
		fmt.Println("has table")
	}
}

func (c *Config_db) migrate_User() {

	User_login := models.User_Login{}
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

	User_bank := models.User_Bank{}
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

	uh := models.User_History{}
	if !c.DB.HasTable(uh) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&uh) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data user history")
	} else {
		// c.DB.AutoMigrate(&User) // สร้าง table, field ต่างๆที่ไม่เคยมี
	}
}

func (c *Config_db) migrate_other() {

	Otp_history := []models.Otp_History{}
	if !c.DB.HasTable(Otp_history) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Otp_history) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data Otp_history")
	}
}
