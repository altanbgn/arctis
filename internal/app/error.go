package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/altanbgn/arctis/internal/models"
)

func errorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Logger().Error(err)

			return c.JSON(http.StatusNotFound, models.Response{
				Success: false,
				Message: "Not Found",
			})
		}

		return nil
	}
}
