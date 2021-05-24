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

	RegisterHandler := RegisterHandler{c}
	LoginHandler := LoginHandler{c}
	UserHandler := UserHandler{c}
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
		{
			HTTPMethod:     http.MethodPost,
			Endpoint:       "/login",
			HandlerFunc:    LoginHandler.Login,
			MiddlewareFunc: []echo.MiddlewareFunc{},
		},
		{
			HTTPMethod:     http.MethodGet,
			Endpoint:       "/:userid",
			HandlerFunc:    UserHandler.Get_Profile,
			MiddlewareFunc: []echo.MiddlewareFunc{middlewares.AuthMiddleware(c, e)},
		},
	}

	for _, r := range routes {
		e.Add(r.HTTPMethod, "/api/v1"+r.Endpoint, r.HandlerFunc, r.MiddlewareFunc...)
	}
}