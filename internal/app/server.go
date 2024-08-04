package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

  "github.com/altanbgn/arctis/internal/config"
  "github.com/altanbgn/arctis/internal/db"
)

func InitServer() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(errorHandler)
	e.Validator = &ValidatorHandler{Validator: validator.New()}

  config.InitEnv()
  db.InitMongo()
  initRoutes(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	// Start server
	go func() {
		if err := e.Start(fmt.Sprintf(":%s", config.Env.PORT)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal(err)
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
		e.Logger.Fatal("Shutting down the server")
	}

	return e
}
