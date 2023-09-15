package ping

import "github.com/labstack/echo/v4"

type Ping struct {
}

func (pr *PingHandler) RegisterPing(server *echo.Echo) {
	server.GET("/ping", pr.Ping)
}

func (pr *PingHandler) Ping(c echo.Context) error {
	return c.JSON(200, "pong")
}
