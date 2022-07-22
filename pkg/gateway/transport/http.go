package http

import (
	"github.com/agniBit/whatsapp/pkg/gateway"
	whatsapp "github.com/agniBit/whatsapp/pkg/gateway/transport/whatsapp"
	"github.com/labstack/echo/v4"
)

func NewHTTP(
	gatewayS gateway.Service,
	e *echo.Group,
) {
	whatsapp.NewHTTP(gatewayS, e)
}
