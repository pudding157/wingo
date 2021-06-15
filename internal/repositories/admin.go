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

	GetAdminSettingBot() (*models.AdminSettingBotResult, error)
	PostAdminSettingBot(a models.Admin_Bank_Condition) (bool, error)
}

type AdminRepo struct {
	c *app.Config
}

func NewAdminRepo(c *app.Config) *AdminRepo {
	return &AdminRepo{c: c}
}

func (r *AdminRepo) PostHome(pc models.Page_Content) (*models.Page_Content, error) {

	fmt.Println("Post all text in home")
	pc.Id = 1
	if err := r.c.DB.Save(&pc).Error; err != nil {
		fmt.Println("h.DB.Find(&pc) => ", err)
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
		bc.CreatedBy = b.CreatedBy
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

func (r *AdminRepo) GetAdminSettingBot() (*models.AdminSettingBotResult, error) {
	as := models.AdminSettingBotResult{}

	rs, err := r.c.DB.Model(&models.Admin_Bank_Condition{}).Select("id, is_active, price_start, price_end, bank_id, account_number").Rows()
	fmt.Println(rs)
	if err != nil {
		fmt.Println("err DB.Find(&Admin_Bank_Condition => ", err)
		return nil, err
	}
	for rs.Next() {
		r.c.DB.ScanRows(rs, &as)
		rs.Close()
	}
	fmt.Println("as => ", as)

	return &as, nil
}

// use AdminBank
func (r *AdminRepo) PostAdminSettingBot(a models.Admin_Bank_Condition) (bool, error) {
	fmt.Println("Post Admin system", a)
	_now := time.Now().UTC()

	as := &models.Admin_Setting{}
	err := r.c.DB.Where("is_active = true").Last(&as).Error
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Setting) => ", err)
		return false, err
	}

	// as.DepositWithdraw = a.DepositWithdraw
	// as.Bet = a.Bet
	// as.CancelBet = a.CancelBet

	as.UpdatedAt = _now
	as.UpdatedBy = r.c.UI

	// if err := r.c.DB.Save(&as).Error; err != nil {
	// 	fmt.Println("h.DB.Find(&as) => ", err)
	// 	return false, err
	// }
	// fmt.Println("h.DB.save Admin system", as)

	return true, nil
}
