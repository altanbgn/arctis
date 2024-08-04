package account

import (
	"net/http"
  "time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/altanbgn/arctis/internal/models"
)

func GetByUsernameHandler(c echo.Context) error {
	ctx := c.Request().Context()

	foundUser, err := GetByUsernameService(ctx, c.Param("username"))
  if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
      Success: false,
      Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response{
    Success: true,
    Data:    foundUser,
	})
}

func GetByIDHandler(c echo.Context) error {
  ctx := c.Request().Context()

  id, err := primitive.ObjectIDFromHex(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  foundUser, err := GetByIDService(ctx, id)
  if err != nil {
    return c.JSON(http.StatusNotFound, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  return c.JSON(http.StatusOK, models.Response{
    Success: true,
    Data:    foundUser,
  })
}

func CreateHandler(c echo.Context) error {
  r := CreateAccountPayload{
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
  }

  if err := c.Bind(&r); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  if err := c.Validate(&r); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  if err := CreateService(c.Request().Context(), r); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  return c.JSON(http.StatusAccepted, models.Response{
    Success: true,
  })
}

func UpdateHandler(c echo.Context) error {
  r := UpdateAccountPayload{
    UpdatedAt: time.Now(),
  }

  id, err := primitive.ObjectIDFromHex(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  if err := c.Bind(&r); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  if err := c.Validate(&r); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }


  if err := UpdateService(c.Request().Context(), id, r); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  return c.JSON(http.StatusAccepted, models.Response{
    Success: true,
  })
}

func DeleteHandler(c echo.Context) error {
  id, err := primitive.ObjectIDFromHex(c.Param("id"))
  if err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  if err := DeleteService(c.Request().Context(), id); err != nil {
    return c.JSON(http.StatusBadRequest, models.Response{
      Success: false,
      Message: err.Error(),
    })
  }

  return c.JSON(http.StatusAccepted, models.Response{
    Success: true,
  })
}
