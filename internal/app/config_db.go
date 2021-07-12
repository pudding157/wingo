package app

import (
	"fmt"
	"time"
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

	Admin_Bank := []models.AdminBank{
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
	_now := time.Now().UTC()

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
		c.DB.AutoMigrate(&User) // สร้าง table, field ต่างๆที่ไม่เคยมี
		// ถ้าจะเพิ่ม unique ถ้า ในตารางมีข้อมูลซ้ำจะไม่สามารถทำได้
		// c.DB.Model(&User).AddUniqueIndex("affiliate", "affiliate")
	}

	User_bank := models.User_Bank{}
	if !c.DB.HasTable(User_bank) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&User_bank) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data User_bank")
	} else {
		// ถ้าจะเพิ่ม unique ถ้า ในตารางมีข้อมูลซ้ำจะไม่สามารถทำได้
		// c.DB.Model(&User_bank).AddUniqueIndex("bank_account", "bank_account")

	}

	User_Wallet := models.User_Wallet{}
	if !c.DB.HasTable(User_Wallet) {
		// fmt.Println("No table")
		c.DB.AutoMigrate(&User_Wallet) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data User_Wallet")
		cc := 0
		c.DB.Model(&User).Count(&cc).Find(&User)
		if cc > 0 {
			for _, _u := range User {
				uw := models.User_Wallet{}
				uw.UserId = _u.Id
				uw.Amount = 0
				uw.CreatedAt = _now
				uw.UpdatedAt = _now

				if err := c.DB.Save(&uw).Error; err != nil {
					fmt.Println("err uw => ", err)
				}
			}
		}
	} else {
		// c.DB.AutoMigrate(&User_Wallet) // สร้าง table, field ต่างๆที่ไม่เคยมี
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
	_now := time.Now().UTC()

	Otp_history := []models.Otp_History{}
	if !c.DB.HasTable(Otp_history) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Otp_history) // สร้าง table, field ต่างๆที่ไม่เคยมี
		fmt.Println("migrate data Otp_history")
	}

	Page_Content := models.Page_Content{}
	if !c.DB.HasTable(Page_Content) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Page_Content) // สร้าง table, field ต่างๆที่ไม่เคยมี
		Page_Content.Id = 1
		Page_Content.CreatedAt = _now
		Page_Content.UpdatedAt = _now
		Page_Content.RunningText = "start running text"
		c.DB.Create(Page_Content)
		fmt.Println("migrate data Page_Content")
	}

	Blog_Content := models.Blog_Content{}
	if !c.DB.HasTable(Blog_Content) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Blog_Content) // สร้าง table, field ต่างๆที่ไม่เคยมี
		Blog_Content.Id = 1
		Blog_Content.Title = "title 1"
		Blog_Content.Content = "<h1>Success!</h1><br>This content has been entered into database.<br>"
		Blog_Content.IsActive = true
		Blog_Content.CreatedBy = 1
		Blog_Content.UpdatedBy = 1
		Blog_Content.CreatedAt = _now
		Blog_Content.UpdatedAt = _now
		c.DB.Create(Blog_Content)
		fmt.Println("migrate data Blog_Content")
	}

	Admin_Setting := models.Admin_Setting{}
	if !c.DB.HasTable(Admin_Setting) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Admin_Setting) // สร้าง table, field ต่างๆที่ไม่เคยมี
		Admin_Setting.Id = 1

		Admin_Setting.DepositWithdraw = true
		Admin_Setting.Bet = true
		Admin_Setting.CancelBet = true
		Admin_Setting.IsActive = true

		Admin_Setting.CreatedBy = 1
		Admin_Setting.UpdatedBy = 1
		Admin_Setting.CreatedAt = _now
		Admin_Setting.UpdatedAt = _now
		c.DB.Create(Admin_Setting)
		fmt.Println("migrate data Admin_Setting")
	} else {
		// c.DB.AutoMigrate(&Admin_Setting)
	}

	Admin_Bank_Condition := models.Admin_Bank_Condition{}
	if !c.DB.HasTable(Admin_Bank_Condition) {
		fmt.Println("No table")
		c.DB.AutoMigrate(&Admin_Bank_Condition) // สร้าง table, field ต่างๆที่ไม่เคยมี
		// Admin_Bank_Condition.Id = 1
		// Admin_Bank_Condition.BankId = 4
		// Admin_Bank_Condition.BankAccount = "6382625487"
		// Admin_Bank_Condition.PriceStart = 0.00
		// Admin_Bank_Condition.PriceEnd = 1000
		// Admin_Bank_Condition.IsActive = true
		// Admin_Bank_Condition.CreatedBy = 1
		// Admin_Bank_Condition.UpdatedBy = 1
		// Admin_Bank_Condition.CreatedAt = _now
		// Admin_Bank_Condition.UpdatedAt = _now
		// fmt.Println("_now", Admin_Bank_Condition)
		// c.DB.Create(Admin_Bank_Condition)
		fmt.Println("migrate data Admin_Bank_Condition")
	}
}
