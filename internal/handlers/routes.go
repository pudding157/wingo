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

	BankRepo := repositories.NewBankRepo(c)
	BankHandler := NewBankHandler(BankRepo)

	HomeRepo := repositories.NewHomeRepo(c)
	HomeHandler := NewHomeHandler(HomeRepo)

	AdminRepo := repositories.NewAdminRepo(c)
	AdminHandler := NewAdminHandler(AdminRepo)

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
		{ // this sprint
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
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/bank/admin",
			HandlerFunc:    BankHandler.GetAdminBanks,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/login",
			HandlerFunc:    LoginHandler.Login,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/logout",
			HandlerFunc:    LoginHandler.Logout,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/home",
			HandlerFunc:    HomeHandler.GetHomeDetail,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/blog",
			HandlerFunc:    HomeHandler.GetBlogs,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/blog/:id",
			HandlerFunc:    AdminHandler.GetBlog,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/user/profile",
			HandlerFunc:    UserHandler.GetProfile,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/user/change-password",
			HandlerFunc:    UserHandler.ChangePassword,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/user/payment/deposit",
			HandlerFunc:    PaymentHandler.Deposit,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/user/payment/withdraw",
			HandlerFunc:    PaymentHandler.Withdraw,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/user/payment/transactions", // /:type = /all | /withdraw | /deposit
			HandlerFunc:    PaymentHandler.Transactions,
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
	AddRoutesAdmin(e, c, *HomeHandler, *AdminHandler)
	return nil
}

func AddRoutesAdmin(e *echo.Echo, c *app.Config, HomeHandler HomeHandler, AdminHandler AdminHandler) {
	routes := []route{

		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/admin/blog",
			HandlerFunc:    HomeHandler.GetBlogs,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/admin/blog/:id",
			HandlerFunc:    AdminHandler.GetBlog,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/admin/home",
			HandlerFunc:    AdminHandler.PostHome,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/admin/blog",
			HandlerFunc:    AdminHandler.PostBlog,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/admin/wallet",
			HandlerFunc:    AdminHandler.GetWallets,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/admin/setting",
			HandlerFunc:    AdminHandler.GetAdminSettingSystem,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/admin/setting",
			HandlerFunc:    AdminHandler.PostAdminSettingSystem,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/admin/setting/bot",
			HandlerFunc:    AdminHandler.GetAdminSettingBots,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
		{ // this sprint
			HTTPMethod:     http.MethodDelete,
			Endpoint:       "/admin/setting/bot",
			HandlerFunc:    AdminHandler.DeleteAdminSettingBot,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthAdminMiddleware(c, e)},
		},
	}
	for _, r := range routes {

		e.Add(r.HTTPMethod, "/api/v1"+r.Endpoint, r.HandlerFunc, r.MiddlewareFunc...)
	}
}
