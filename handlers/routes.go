package handlers

import (
	// "winapp/internal/app"
	// "winapp/internal/middlewares"
	// "winapp/internal/repositories"
	// "winapp/internal/services"
	"winapp/app"
	// "winapp/middlewares"

	// "github.com/labstack/echo/v4/middleware"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type route struct {
	HTTPMethod     string
	Endpoint       string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

// NewRouter func
func NewRouter(e *echo.Echo, db *gorm.DB, c *app.Config) error {
	register_module(e, db, c)
	// merchantRepo := repositories.NewMerchantRepo(c)
	// merchantService := services.NewMerchantService(merchantRepo)
	// productRepo := repositories.NewProductRepo(c)
	// productService := services.NewProductService(productRepo)
	// reportService := services.NewReportService()
	// merchantHandler := NewMerchantHandler(merchantService, productService, reportService)
	// routes := []route{
	// 	{
	// 		HTTPMethod:     http.MethodGet,
	// 		Endpoint:       "/merchant/information/:id",
	// 		HandlerFunc:    merchantHandler.GetMerchantByID,
	// 		MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
	// 	},
	// 	{
	// 		HTTPMethod:     http.MethodPost,
	// 		Endpoint:       "/merchant/register",
	// 		HandlerFunc:    merchantHandler.Register,
	// 		MiddlewareFunc: []echo.MiddlewareFunc{},
	// 	},
	// 	{
	// 		HTTPMethod:     http.MethodPost,
	// 		Endpoint:       "/merchant/update",
	// 		HandlerFunc:    merchantHandler.Update,
	// 		MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
	// 	},
	// 	{
	// 		HTTPMethod:     http.MethodPost,
	// 		Endpoint:       "/merchant/:id/product",
	// 		HandlerFunc:    merchantHandler.CreateProduct,
	// 		MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
	// 	},
	// 	{
	// 		HTTPMethod:     http.MethodGet,
	// 		Endpoint:       "/merchant/:id/products",
	// 		HandlerFunc:    merchantHandler.GetProducts,
	// 		MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
	// 	},
	// 	{
	// 		HTTPMethod:     http.MethodGet,
	// 		Endpoint:       "/merchant/:id/report",
	// 		HandlerFunc:    merchantHandler.GenReport,
	// 		MiddlewareFunc: []echo.MiddlewareFunc{middlewares.RequestHandlerMiddleware(c)},
	// 	},
	// }

	// for _, r := range routes {
	// 	e.Add(r.HTTPMethod, "/api"+r.Endpoint, r.HandlerFunc, r.MiddlewareFunc...)
	// }
	return nil
}

func register_module(e *echo.Echo, db *gorm.DB, c *app.Config) {

	RegisterHandler := RegisterHandler(db)
	BankHandler := BankHandler(db)
	LoginHandler := LoginHandler(db)

	routes := []route{
		{
			HTTPMethod: http.MethodGet,
			Endpoint:   "/",
			HandlerFunc: func(c echo.Context) error {
				return c.String(http.StatusOK, "Hello, World!")
			},
			// MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/register/otp/send",
			HandlerFunc:    RegisterHandler.Otp_send,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/register/otp",
			HandlerFunc:    RegisterHandler.Otp,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/register",
			HandlerFunc:    RegisterHandler.Register,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/bank",
			HandlerFunc:    BankHandler.Get_all_bank,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/login",
			HandlerFunc:    LoginHandler.Login,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
	}
	// // get all bank
	// e.GET("/api/v1/bank", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "get all bank")
	// })

	// // send otp
	// e.POST("/api/v1/register/otp/send", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "send otp")
	// })

	// // post check otp
	// e.POST("/api/v1/register/otp", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "post check otp")
	// })

	for _, r := range routes {
		e.Add(r.HTTPMethod, "/api/v1"+r.Endpoint, r.HandlerFunc, r.MiddlewareFunc...)
	}
}
