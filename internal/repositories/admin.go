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

	GetAdminSettingBot() (*models.AdminSettingBotListResult, error)
	PostAdminSettingBot(a models.AdminSettingBotListBind) (bool, error)
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

func (r *AdminRepo) GetAdminSettingBot() (*models.AdminSettingBotListResult, error) {
	as := []models.AdminSettingBotResult{}

	a := &models.Admin_Setting{}
	err := r.c.DB.Where("is_active = true").Find(&a).Error
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Setting) => ", err)
		return nil, err
	}

	s := []models.Admin_Bank_Condition{}
	err = r.c.DB.Find(&s).Error //.Where("is_active = true")
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

	ar := models.AdminSettingBotListResult{
		IsBotActive:           a.IsBotActive,
		AdminSettingBotResult: as,
	}

	fmt.Println("as => ", as)

	return &ar, nil
}

// use AdminBank
func (r *AdminRepo) PostAdminSettingBot(a models.AdminSettingBotListBind) (bool, error) {
	fmt.Println("Post Admin setting bot", a)

	for _, ab := range a.Admin_Bank_Condition {
		if ab.PriceStart >= ab.PriceEnd {
			return false, errors.New("price start has higher or equal then price end")
		}
	}
	_now := time.Now().UTC()

	ast := &models.Admin_Setting{}
	err := r.c.DB.Where("is_active = true").Find(&ast).Error
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Setting) => ", err)
		return false, err
	}
	ast.IsBotActive = a.IsBotActive
	ast.UpdatedBy = r.c.UI
	ast.UpdatedAt = _now

	as := []models.Admin_Bank_Condition{}
	err = r.c.DB.Find(&as).Error //.Where("is_active = true")
	if err != nil {
		fmt.Println("h.DB.Find(&Admin_Bank_Condition) => ", err)
		return false, err
	}
	sp := []float64{} // start price
	se := []float64{} // start end
	for i, ab := range a.Admin_Bank_Condition {
		fmt.Println("Current PriceStart => ", ab.PriceStart)
		fmt.Println("Current PriceEnd => ", ab.PriceEnd)
		if ab.Id == 0 {
			a.Admin_Bank_Condition[i].CreatedBy = r.c.UI
			a.Admin_Bank_Condition[i].CreatedAt = _now
		} else {
			for _, b := range as {
				if b.Id == ab.Id {
					// Found!
					a.Admin_Bank_Condition[i].CreatedBy = b.CreatedBy
					a.Admin_Bank_Condition[i].CreatedAt = b.CreatedAt
					a.Admin_Bank_Condition[i].AccessToken = b.AccessToken
					a.Admin_Bank_Condition[i].ApiRefresh = b.ApiRefresh
					a.Admin_Bank_Condition[i].DeviceId = b.DeviceId
				}
			}
		}
		a.Admin_Bank_Condition[i].UpdatedBy = r.c.UI
		a.Admin_Bank_Condition[i].UpdatedAt = _now
		a.Admin_Bank_Condition[i].IsActive = true
		// a.Admin_Bank_Condition[i].IsActive = a.IsBotActive

		if i != 0 {
			for u := range sp {
				if (ab.PriceStart <= se[u] && ab.PriceStart >= sp[u]) ||
					(ab.PriceEnd >= sp[u] && ab.PriceEnd <= se[u]) {
					fmt.Println("ติด gap price")
					return false, errors.New("price start or price end had value in between values before.")
				} else {
					fmt.Println("pass")
				}
			}
		}
		sp = append(sp, ab.PriceStart) // start price
		se = append(se, ab.PriceEnd)   // start end
	}

	tx := r.c.DB.Begin()
	ids := []int{}
	for _, bc := range a.Admin_Bank_Condition {
		if err := tx.Save(&bc).Error; err != nil {
			fmt.Println("h.DB.Find(&Admin_Bank_Condition) => ", err)
			tx.Rollback()
			return false, err
		}
		ids = append(ids, bc.Id)
	}
	c := []models.Admin_Bank_Condition{}
	// sub := tx.Find(&c, ids).SubQuery()
	if err := tx.Not(ids).Find(&c).Error; err != nil {
		fmt.Println("id NOT IN ? => ", err)
		tx.Rollback()
		return false, err
	} else {
		for _, bc := range c {
			bc.IsActive = false
			bc.UpdatedBy = r.c.UI
			bc.UpdatedAt = _now
			if err := tx.Save(&bc).Error; err != nil {
				fmt.Println("h.DB.Find(&c) => ", err)
				tx.Rollback()
				return false, err
			}
		}
	}

	if err := tx.Save(&ast).Error; err != nil {
		fmt.Println("h.DB.Find(&Admin_Setting) => ", err)
		tx.Rollback()
		return false, err
	}

	tx.Commit()

	fmt.Println("h.DB.save Admin_Bank_Condition", a)

	return true, nil
}

// func RemoveIndex(s []models.Admin_Bank_Condition, index int) []models.Admin_Bank_Condition {
// 	return append(s[:index], s[index+1:]...)
// }
func (r *AdminRepo) GetBlog(id int) (*models.Blog_Content, error) {

	bc := models.Blog_Content{}

	if err := r.c.DB.Find(&bc, id).Error; err != nil {
		fmt.Println("h.DB.Find Blog_Content error => ", err)
		return nil, err
	}
	fmt.Println("h.DB.Find Blog Content", bc)
	return &bc, nil
}
