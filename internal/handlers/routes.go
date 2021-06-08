package handlers

import (
	// "winapp/internalinternal/app"
	// "winapp/internalinternal/middlewares"
	// "winapp/internalinternal/repositories"
	// "winapp/internalinternal/services"

	// "winapp/internalmiddlewares"
	"net/http"
	"winapp/internal/app"
	"winapp/internal/middlewares"
	"winapp/internal/repositories"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
	// _ "github.com/swaggo/echo-swagger/example/docs" // docs is generated by Swag CLI, you have to import it.
	_ "winapp/docs"
)

type route struct {
	HTTPMethod     string
	Endpoint       string
	HandlerFunc    echo.HandlerFunc
	MiddlewareFunc []echo.MiddlewareFunc
}

// NewRouter func
func NewRouter(e *echo.Echo, c *app.Config) error {
	register_module(e, c)
	return nil
}

func register_module(e *echo.Echo, c *app.Config) {

	BankRepo := repositories.NewBankRepo(c)
	BankHandler := NewBankHandler(BankRepo)

	LoginRepo := repositories.NewLoginRepo(c)
	LoginHandler := NewLoginHandler(LoginRepo)

	UserRepo := repositories.NewUserRepo(c, LoginRepo)
	UserHandler := NewUserHandler(UserRepo)

	RegisterRepo := repositories.NewRegisterRepo(c, LoginRepo)
	RegisterHandler := NewRegisterHandler(RegisterRepo)

	PaymentRepo := repositories.NewPaymentRepo(c)
	PaymentHandler := NewPaymentHandler(PaymentRepo)
	// RegisterHandler := RegisterHandler{c}
	// LoginHandler := LoginHandler{c}
	// UserHandler := UserHandler{c}
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
			HandlerFunc:    BankHandler.GetBanks,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/bank/admin",
			HandlerFunc:    BankHandler.GetAdminBanks, // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/login",
			HandlerFunc:    LoginHandler.Login,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{ // this sprint
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/logout",
			HandlerFunc:    LoginHandler.Logout, // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/user/profile",
			HandlerFunc:    UserHandler.GetProfile, // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/user/change-password",
			HandlerFunc:    UserHandler.ChangePassword, // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/user/payment/deposit",
			HandlerFunc:    PaymentHandler.Deposit, // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/user/payment/withdraw",
			HandlerFunc:    PaymentHandler.Withdraw, // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/user/payment/transactions", // /:type = /all | /withdraw | /deposit
			HandlerFunc:    PaymentHandler.Transactions,  // success
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/user/affiliate",
			HandlerFunc:    UserHandler.GetAffiliate,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
	}
	if c.Env == "dev" {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}
	for _, r := range routes {
		e.Add(r.HTTPMethod, "/api/v1"+r.Endpoint, r.HandlerFunc, r.MiddlewareFunc...)
	}
}
