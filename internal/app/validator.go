package app

import (
  "net/http"

  "github.com/go-playground/validator/v10"
  "github.com/labstack/echo/v4"
)

type ValidatorHandler struct {
  Validator *validator.Validate
}

func (cv *ValidatorHandler) Validate(i interface{}) error {
  if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }

  return nil
}
