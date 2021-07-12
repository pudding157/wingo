package repositories

import (
	"fmt"
	"winapp/internal/app"
	"winapp/internal/models"
)

type HomeRepository interface {
	GetHomeDetail() (*models.Page_Content, error)
	GetBlogs(t models.LoadMoreModel) (*[]models.Blog_Content, int64, error)
}

type HomeRepo struct {
	c *app.Config
}

func NewHomeRepo(c *app.Config) *HomeRepo {
	return &HomeRepo{c: c}
}

func (r *HomeRepo) GetHomeDetail() (*models.Page_Content, error) {

	fmt.Println("Get all text in home")

	pc := models.Page_Content{}

	if err := r.c.DB.Find(&pc, "id = 1").Error; err != nil {
		fmt.Println("h.DB.Find(&pc) => ", err)
		return nil, err
	}
	fmt.Println("h.DB.Find page content", pc)
	return &pc, nil
}

func (r *HomeRepo) GetBlogs(t models.LoadMoreModel) (*[]models.Blog_Content, int64, error) {
	fmt.Println("Get all Blogs in home")
	fmt.Println("transaction type => ", t)

	bc := []models.Blog_Content{}
	c := 0
	cs := "is_active = 1"
	if err := r.c.DB.Model(&bc).Where(cs).Count(&c).Limit(t.Take).Offset(t.Skip).Order("created_at desc").Find(&bc, cs).Error; err != nil {
		fmt.Println("h.DB.Find Blog_Content error => ", err)
		return nil, int64(0), err
	}
	fmt.Println("h.DB.Find page content", bc)
	return &bc, int64(c), nil
}
