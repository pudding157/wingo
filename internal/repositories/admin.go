package repositories

import (
	"fmt"
	"time"
	"winapp/internal/app"
	"winapp/internal/models"
)

type AdminRepository interface {
	PostHome(pc models.Page_Content) (*models.Page_Content, error)
	PostBlog(bc models.Blog_Content) (*models.Blog_Content, error)
	GetWallets() (*models.WalletsResult, error)
}

type AdminRepo struct {
	c *app.Config
}

func NewAdminRepo(c *app.Config) *AdminRepo {
	return &AdminRepo{c: c}
}

func (r *AdminRepo) PostHome(pc models.Page_Content) (*models.Page_Content, error) {

	fmt.Println("Post all text in home")

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
