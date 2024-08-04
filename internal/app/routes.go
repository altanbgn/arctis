package app

import (
  "net/http"

  "github.com/labstack/echo/v4"
  "github.com/altanbgn/arctis/internal/domains/account"
)

func initRoutes(e *echo.Echo) {
  v1 := e.Group("/v1")

  v1.GET("/health", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{
      "message": "OK",
    })
  })

  v1.POST("/soap", soapHandler)

  account.RegisterAccountRoutes(v1)
}
