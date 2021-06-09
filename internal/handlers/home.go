package handlers

import (
	"fmt"
	"strconv"
	"winapp/internal/models"
	"winapp/internal/repositories"

	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type HomeHandler struct {
	// // DB *gorm.DB
	// // R  *redis.Client
	// c *app.Config
	Repo repositories.HomeRepository
}

func NewHomeHandler(repo repositories.HomeRepository) *HomeHandler {
	return &HomeHandler{Repo: repo}
}

func (r *HomeHandler) GetHomeDetail(c echo.Context) error {

	fmt.Println("Get all home details")

	h, err := r.Repo.GetHomeDetail()

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = h

	return c.JSON(http.StatusOK, _res)
}
func (r *HomeHandler) PostHome(c echo.Context) error {

	fmt.Println("Post all home details")

	pc := &models.Page_Content{}
	c.Bind(&pc)
	h, err := r.Repo.PostHome(*pc)

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = h

	return c.JSON(http.StatusOK, _res)
}

func (r *HomeHandler) GetBlogs(c echo.Context) error {

	fmt.Println("Get all home details")

	qs := c.QueryParams()
	fmt.Println("QueryParams => ", qs)

	l := models.LoadMoreModel{}

	// l.Type = qs.Get("type")

	// if l.Type != "all" && l.Type != "deposit" && l.Type != "withdraw" {
	// 	l.Type = "all"
	// }
	s, err := strconv.Atoi(qs.Get("skip"))
	if err != nil {
		s = 0
	}
	l.Skip = s

	t, err := strconv.Atoi(qs.Get("take"))
	if err != nil {
		t = 20
	}
	l.Take = t

	h, err := r.Repo.GetBlogs(l)

	if err != nil {
		_res := models.ErrorResponse{}
		_res.Error = "Validation Failed"
		_res.ErrorMessage = err.Error()
		// _res.Error_message = [{"phone_number": "phone number must be at least 10 digits."}]
		_res.Error_code = strconv.Itoa(http.StatusInternalServerError)
		return c.JSON(http.StatusInternalServerError, _res)
	}

	_res := models.Response{}
	_res.Data = h

	return c.JSON(http.StatusOK, _res)
}
