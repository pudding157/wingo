package repositories

import (
	"errors"
	"fmt"
	"time"
	"winapp/internal/app"
	"winapp/internal/models"
)

type AdminRepository interface {
	PostHome(pc models.Page_Content) (*models.Page_Content, error)
	PostBlog(bc models.Blog_Content) (*models.Blog_Content, error)
	GetWallets() (*models.WalletsResult, error)

	GetAdminSettingSystem() (*models.AdminSettingSystemResult, error)
	PostAdminSettingSystem(a models.Admin_Setting) (bool, error)

	GetAdminSettingBot() (*[]models.AdminSettingBotResult, error)
	PostAdminSettingBot(a models.Admin_Bank_Condition) (bool, error)
	GetBlog(id int) (*models.Blog_Content, error)
}

type AdminRepo struct {
	c *app.Config
}

func NewAdminRepo(c *app.Config) *AdminRepo {
	return &AdminRepo{c: c}
}

func (r *AdminRepo) PostHome(pc models.Page_Content) (*models.Page_Content, error) {

	fmt.Println("Post all text in home")
	as := &models.Page_Content{}
	err := r.c.DB.Find(&as).Error
	if err != nil {
		fmt.Println("h.DB.Find(&Page_Content) => ", err)
		return nil, err
	}

	as.UpdatedAt = time.Now().UTC()
	as.RunningText = pc.RunningText

	if err := r.c.DB.Save(&as).Error; err != nil {
		fmt.Println("h.DB.Find(&as) => ", err)
		return nil, err
	}
	fmt.Println("h.DB.save page content", pc)
	return &pc, nil
}

func (r *AdminRepo) PostBlog(bc models.Blog_Content) (*models.Blog_Content, error) {

	fmt.Println("Post Blog", bc)
	_now := time.Now().UTC()
	if bc.Id == 0 {
		bc.CreatedBy = r.c.UI
		bc.CreatedAt = _now
		bc.IsActive = true
	} else {

		b := &models.Blog_Content{}
		err := r.c.DB.Find(&b, "id = ?", bc.Id).Error
		if err != nil {
			fmt.Println("err DB.Find(&bc => ", err)
			return nil, err
		}
		// bc.CreatedBy = b.CreatedBy
		bc.IsActive = true
		// bc.CreatedAt = b.CreatedAt
	}
	bc.UpdatedBy = r.c.UI
	bc.UpdatedAt = _now
	if err := r.c.DB.Save(&bc).Error; err != nil {
		fmt.Println("h.DB.Find(&bc) => ", err)
		return nil, err
	}
	fmt.Println("h.DB.save blog content", bc)
	return &bc, nil
}

func (r *AdminRepo) GetWallets() (*models.WalletsResult, error) {
	wr := models.WalletsResult{}
	wr.Amount = 0
	wr.Count = 0
	rs, _ := r.c.DB.Model(&models.User_Wallet{}).Select("sum(amount) as amount, COUNT(amount) as count").Rows()

	fmt.Println(rs)
	for rs.Next() {
		r.c.DB.ScanRows(rs, &wr)
		rs.Close()
	}
	fmt.Println("wr => ", wr)
	// err := r.c.DB.Model(&models.User_Wallet{}).Error
	// if err != nil {
	// 	fmt.Println("err DB.Find(&User_Wallet => ", err)
	// 	return nil, err
	// }

	return &wr, nil
}

func (r *AdminRepo) GetAdminSettingSystem() (*models.AdminSettingSystemResult, error) {
	as := models.AdminSettingSystemResult{}

	rs, err := r.c.DB.Model(&models.Admin_Setting{}).Where("is_active = true").Select("id, deposit_withdraw, bet, cancel_bet").Rows()
	fmt.Println(rs)
	if err != nil {
		fmt.Println("err DB.Find(&User_Wallet => ", err)
		return nil, err
	}
	for rs.Next() {
		r.c.DB.ScanRows(rs, &as)
		rs.Close()
	}
	if as.Id == 0 {
		return nil, errors.New("not found row data.")
	}
	fmt.Println("as => ", as)

	return &as, nil
}

// use AdminBank
func (r *AdminRepo) PostAdminSettingSystem(a models.Admin_Setting) (bool, error) {
	fmt.Println("Post Admin system", a)
	_now := time.Now().UTC()

	as := &models.Admin_Setting{}
	err := r.c.DB.Where("is_active = true").Last(&as).Error
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Setting) => ", err)
		return false, err
	}

	as.DepositWithdraw = a.DepositWithdraw
	as.Bet = a.Bet
	as.CancelBet = a.CancelBet

	as.UpdatedAt = _now
	as.UpdatedBy = r.c.UI

	if err := r.c.DB.Save(&as).Error; err != nil {
		fmt.Println("h.DB.Find(&as) => ", err)
		return false, err
	}
	fmt.Println("h.DB.save Admin system", as)

	return true, nil
}

func (r *AdminRepo) GetAdminSettingBot() (*[]models.AdminSettingBotResult, error) {
	as := []models.AdminSettingBotResult{}

	// rs, err := r.c.DB.Model(&models.Admin_Bank_Condition{}).Select("id, is_active, price_start, price_end, bank_id, bank_account").Rows()
	// fmt.Println("rs => ", rs)
	// if err != nil {
	// 	fmt.Println("err DB.Find(&Admin_Bank_Condition => ", err)
	// 	return nil, err
	// }
	// for rs.Next() {
	// 	r.c.DB.ScanRows(rs, &as)
	// 	rs.Close()
	// }
	s := []models.Admin_Bank_Condition{}
	err := r.c.DB.Find(&s).Error //.Where("is_active = true")
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Bank_Condition) => ", err)
		return nil, err
	}
	if len(s) > 0 {
		for _, r := range s {
			b := models.AdminSettingBotResult{}
			b.Id = r.Id
			b.IsActive = r.IsActive
			b.PriceStart = r.PriceStart
			b.PriceEnd = r.PriceEnd
			b.BankId = r.BankId
			b.BankAccount = r.BankAccount
			as = append(as, b)
		}
	} else {
		return nil, errors.New("not found row data.")
	}

	fmt.Println("as => ", as)

	return &as, nil
}

// use AdminBank
func (r *AdminRepo) PostAdminSettingBot(a models.Admin_Bank_Condition) (bool, error) {
	fmt.Println("Post Admin setting bot", a)

	if a.PriceStart >= a.PriceEnd {
		return false, errors.New("price start has higher or equal then price end")
	}

	as := []models.Admin_Bank_Condition{}
	err := r.c.DB.Where("is_active = true").Find(&as).Error
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Bank_Condition) => ", err)
		return false, err
	}

	fmt.Println("Current PriceStart => ", a.PriceStart)
	fmt.Println("Current PriceEnd => ", a.PriceEnd)
	for _, abc := range as {
		fmt.Println(abc.Id, ": PriceStart => ", abc.PriceStart)
		fmt.Println(abc.Id, ": PriceEnd => ", abc.PriceEnd)
		if (a.PriceStart <= abc.PriceEnd && a.PriceStart >= abc.PriceStart) ||
			(a.PriceEnd >= abc.PriceStart && a.PriceEnd <= abc.PriceEnd) {
			fmt.Println("ติด gap price")
			return false, errors.New("price start or price end had value in between values before.")
		} else {
			fmt.Println("pass")
		}
	}

	_now := time.Now().UTC()
	if a.Id == 0 {
		a.CreatedBy = r.c.UI
		a.CreatedAt = _now
	} else {
		b := &models.Admin_Bank_Condition{}
		err := r.c.DB.Find(&b, "id = ?", a.Id).Error
		if err != nil {
			fmt.Println("err DB.Find(&bc => ", err)
			return false, err
		}
		a.CreatedBy = b.CreatedBy
		a.CreatedAt = b.CreatedAt
		a.AccessToken = b.AccessToken
		a.ApiRefresh = b.ApiRefresh
		a.DeviceId = b.DeviceId
	}
	a.UpdatedAt = _now
	a.UpdatedBy = r.c.UI

	if err := r.c.DB.Save(&a).Error; err != nil {
		fmt.Println("h.DB.Find(&a) => ", err)
		return false, err
	}
	fmt.Println("h.DB.save Admin_Bank_Condition", a)

	return true, nil
}

func (r *AdminRepo) GetBlog(id int) (*models.Blog_Content, error) {

	bc := models.Blog_Content{}

	if err := r.c.DB.Find(&bc, id).Error; err != nil {
		fmt.Println("h.DB.Find Blog_Content error => ", err)
		return nil, err
	}
	fmt.Println("h.DB.Find Blog Content", bc)
	return &bc, nil
}
