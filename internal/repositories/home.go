package repositories

import (
	"fmt"
	"winapp/internal/app"
	"winapp/internal/models"
)

type HomeRepository interface {
	GetHomeDetail() (*models.Page_Content, error)
	PostHome(pc models.Page_Content) (*models.Page_Content, error)
	GetBlogs(t models.LoadMoreModel) ([]string, error)
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
func (r *HomeRepo) PostHome(pc models.Page_Content) (*models.Page_Content, error) {

	fmt.Println("Get all text in home")

	if err := r.c.DB.Save(&pc).Error; err != nil {
		fmt.Println("h.DB.Find(&pc) => ", err)
		return nil, err
	}
	fmt.Println("h.DB.save page content", pc)
	return &pc, nil
}

func (r *HomeRepo) GetBlogs(t models.LoadMoreModel) ([]string, error) {

	fmt.Println("Get all text in home", t)

	// b := []models.Bank{}

	// // get only isactive
	// if err := r.c.DB.Find(&b, "is_active = 1").Error; err != nil {
	// 	fmt.Println("h.DB.Find(&banks) => ", err)
	// 	return nil, err
	// }
	// fmt.Println("h.DB.Find(&banks) => true =>  ", b)
	return []string{"blog kub"}, nil
}
