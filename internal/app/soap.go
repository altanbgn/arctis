package app

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/altanbgn/arctis/internal/domains/account"
	"github.com/altanbgn/arctis/internal/models"
)

var soapActions = map[string]models.Action{
	// Account
	"accountGetByID":       account.GetByIDAction,
	"accountGetByUsername": account.GetByUsernameAction,
	"accountCreate":        account.CreateAction,
	"accountUpdate":        account.UpdateAction,
	"accountDelete":        account.DeleteAction,
}

func soapHandler(c echo.Context) error {
	defer c.Request().Body.Close()
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to read request")
	}

	soapAction := c.Request().Header.Get("SOAPAction")
	if soapAction == "" {
		return c.String(http.StatusNotFound, "SOAPAction header is missing")
	}

	var envelope models.Envelope
	if err = xml.Unmarshal(body, &envelope); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse XML")
	}

	handler, exists := soapActions[soapAction]
	if !exists {
		return c.String(http.StatusNotFound, "Invalid action request")
	}

	return handler(c, envelope)
}
