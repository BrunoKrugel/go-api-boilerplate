package api

import (
	"context"
	"net/http"

	"github.com/BrunoKrugel/zaplog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Meta struct {
	Server *echo.Echo
}

func NewMeta() *Meta {
	e := echo.New()
	e.HideBanner = true
	// e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/liveness", HealthHandler)
	e.GET("/readiness", EnvironmentHandler)

	return &Meta{
		Server: e,
	}
}

func StartMeta(ctx context.Context, meta *Meta) {
	zaplog.Info(ctx, "Meta server started")

	if err := meta.Server.Start(":8081"); err != nil {
		zaplog.Fatal(ctx, "Failed to start Meta server")
	}
}

func HealthHandler(c echo.Context) error {
	return up(c)
}

func EnvironmentHandler(c echo.Context) error {
	return up(c)
}

func up(c echo.Context) error {
	response := make(map[string]string)
	response["status"] = "UP"
	return c.JSON(http.StatusOK, response)
}
