package handlers

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func InitializeRoutes(e *echo.Echo, db *MyDB) {
	e.GET("/login", db.showLoginPage, ensureNotLoggedIn)
	e.POST("/login", db.performLogin, ensureNotLoggedIn)

	r := e.Group("/logout")
	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte("secret"),
		TokenLookup: "cookie:Authorization",
	}))

	middleware.ErrJWTMissing = echo.NewHTTPError(
		http.StatusBadRequest,
		map[string]string{
			"message": "يجب ان تسجل الدخول",
			"url":     "/login",
		},
	)

	r.GET("", logout)
}
