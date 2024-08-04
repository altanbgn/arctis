package account

import (
	"github.com/labstack/echo/v4"
)

func RegisterAccountRoutes(e *echo.Group) {
	account := e.Group("/account")

	account.GET("/by-username/:username", GetByUsernameHandler)
  account.GET("/by-id/:id", GetByIDHandler)
  account.POST("", CreateHandler)
	account.PUT("/by-id/:id", UpdateHandler)
	account.DELETE("/by-id/:id", DeleteHandler)
}
