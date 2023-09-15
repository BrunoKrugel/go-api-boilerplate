package ping

import (
	"github.com/BrunoKrugel/go-api-boilerplate/internal/business/ping/service"
	"github.com/labstack/echo/v4"
)

type PingHandler struct {
	pingService service.PingServiceInterface
}

func NewPingHandler(pingService service.PingServiceInterface) *PingHandler {
	return &PingHandler{
		pingService: pingService,
	}
}

func (pingHandler *PingHandler) Register(server *echo.Echo) {
	pingHandler.RegisterPing(server)
}
