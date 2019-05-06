package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func ensureNotLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("Authorization")
		if cookie != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "يجب ألا تكون مسجل",
			})
		}
		return next(c)
	}
}
