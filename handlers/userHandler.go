package handlers //เปรียบเสมือน namespace c#

// import (
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// 	"github.com/labstack/echo/v4"
// )

// type UserHandler struct {
// 	DB *gorm.DB
// }

// //ให้เชื่อมต่อฐานข้อมูลเมื่อ Initialize
// func (h *UserHandler) Initialize() {
// 	db, err := gorm.Open("mysql", "root:helloworld@tcp(localhost:6603)/godb?charset=utf8&parseTime=True") //127.0.0.1:3306
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db.AutoMigrate(&User{}) // สร้าง table, field ต่างๆที่ไม่เคยมี

// 	h.DB = db
// }

// type User struct {
// 	Id        uint   `gorm:"primary_key" json:"id"`
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// 	Age       int    `json:"age"`
// 	Email     string `json:"email"`
// }

// func (h *UserHandler) GetAllUser(c echo.Context) error {
// 	users := []User{}

// 	h.DB.Find(&users)

// 	return c.JSON(http.StatusOK, users)
// }

// func (h *UserHandler) GetUser(c echo.Context) error {
// 	id := c.Param("id")
// 	user := User{}

// 	if err := h.DB.Find(&user, id).Error; err != nil {
// 		return c.NoContent(http.StatusNotFound)
// 	}

// 	return c.JSON(http.StatusOK, user)
// }

// func (h *UserHandler) SaveUser(c echo.Context) error {
// 	user := User{}
// 	// if err := c.Bind(&user); err != nil {
// 	// 	return c.NoContent(http.StatusBadRequest)
// 	// }
// 	user.FirstName = c.FormValue("firstName")
// 	user.LastName = c.FormValue("lastName")

// 	_age, _ := strconv.Atoi(c.FormValue("Age"))
// 	user.Age = _age
// 	user.Email = c.FormValue("Email")
// 	log.Print("user1", user)
// 	// if err := h.DB.Save(&user); err != nil {
// 	// 	log.Print("err => ", err)
// 	// 	return c.NoContent(http.StatusInternalServerError)
// 	// }
// 	h.DB.Save(&user)
// 	return c.JSON(http.StatusOK, user)
// }
