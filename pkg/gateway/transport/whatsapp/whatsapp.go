package transport

import (
	"github.com/agniBit/whatsapp/pkg/gateway"
	"github.com/agniBit/whatsapp/util/config"
	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc gateway.Service
	cfg *config.Config
}

func NewHTTP(svc gateway.Service, cfg *config.Config, e *echo.Group) {
	h := &HTTP{svc: svc, cfg: cfg}

	e.GET("/", h.index)
	e.GET("/wa/webhook", h.verifyCallback)

	// event webhook
	e.POST("/wa/webhook", h.messageEventWebhook)
}


func (h *HTTP) index(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"message": "Hello, World!",
	})
}
