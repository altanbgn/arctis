package account

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/altanbgn/arctis/internal/models"
)

func GetByUsernameAction(c echo.Context, envelope models.Envelope) error {
	ctx := c.Request().Context()
  fmt.Println(envelope)

	foundUser, err := GetByUsernameService(ctx, c.Param("username"))
	if err != nil {
    return c.XML(http.StatusNotFound, models.Envelope{
      Namespaces: map[string]string{"soap": "http://schemas.xmlsoap.org/soap/envelope/"},
      Body: models.Body{Fault: models.Fault{String: err.Error()}},
    })
	}

	r := models.Envelope{
    Namespaces: map[string]string{"soap": "http://schemas.xmlsoap.org/soap/envelope/"},
		Body: models.Body{Content: foundUser},
	}

	output, err := xml.MarshalIndent(r, "", "  ")
	if err != nil {
    return c.XML(http.StatusNotFound, models.Envelope{
      Namespaces: map[string]string{"soap": "http://schemas.xmlsoap.org/soap/envelope/"},
      Body: models.Body{Fault: models.Fault{String: err.Error()}},
    })
	}

	return c.XML(http.StatusOK, output)
}

func GetByIDAction(c echo.Context, envelope models.Envelope) error {
	return c.String(http.StatusNotFound, "getbyid")
}

func CreateAction(c echo.Context, envelope models.Envelope) error {
	return c.String(http.StatusNotFound, "create")
}

func UpdateAction(c echo.Context, envelope models.Envelope) error {
	return c.String(http.StatusNotFound, "update")
}

func DeleteAction(c echo.Context, envelope models.Envelope) error {
	return c.String(http.StatusNotFound, "delete")
}
