package transport

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *HTTP) verifyCallback(c echo.Context) error {
	challenge := c.QueryParam("hub.challenge")
	fmt.Println(challenge)
	return c.String(200, challenge)
}

func (h *HTTP) messageEventWebhook(c echo.Context) error {
	return c.JSON(200, http.StatusOK)
}
