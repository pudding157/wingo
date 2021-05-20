package handlers

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Handler struct
type Handler struct {
	DB *gorm.DB
	// mService services.MerchantService
	// pService services.ProductService
	// rService services.ReportService
}

func Initialize() *Handler {

	// root:helloworld@tcp(localhost:6603)/godb?charset=utf8&parseTime=True
	// root:helloworld@tcp(db:3306)/godb?charset=utf8&parseTime=True
	db, err := gorm.Open("mysql", "root:helloworld@tcp(localhost:6603)/godb?charset=utf8&parseTime=True") //127.0.0.1:3306
	if err != nil {
		log.Fatal(err)
	}
	return &Handler{DB: db}
}
