package api

import (
	"context"
	"net/http"

	"github.com/BrunoKrugel/zaplog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTP struct {
	Server *echo.Echo
}

func NewHTTP() *HTTP {
	e := echo.New()
	e.HideBanner = true
	// e.HidePort = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return &HTTP{
		Server: e,
	}
}

func StartHttp(ctx context.Context, http *HTTP) {
	zaplog.Info(ctx, "HTTP server started")

	if err := http.Server.Start(":8080"); err != nil {
		zaplog.Fatal(ctx, "Failed to start HTTP server")
	}
}
