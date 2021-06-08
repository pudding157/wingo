package repositories

import (
	"fmt"
	"winapp/internal/app"
	"winapp/internal/models"
)

type HomeRepository interface {
	GetHomeDetail() ([]string, error)
	GetBlogs(t models.LoadMoreModel) ([]string, error)
}

type HomeRepo struct {
	c *app.Config
}

func NewHomeRepo(c *app.Config) *HomeRepo {
	return &HomeRepo{c: c}
}

func (r *HomeRepo) GetHomeDetail() ([]string, error) {

	fmt.Println("Get all text in home")

	// b := []models.Bank{}

	// // get only isactive
	// if err := r.c.DB.Find(&b, "is_active = 1").Error; err != nil {
	// 	fmt.Println("h.DB.Find(&banks) => ", err)
	// 	return nil, err
	// }
	// fmt.Println("h.DB.Find(&banks) => true =>  ", b)
	return []string{"เลขวิ่ง"}, nil
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
