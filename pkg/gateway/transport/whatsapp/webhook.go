package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/agniBit/whatsapp/type/whatsapp"
	"github.com/labstack/echo/v4"
)

func (h *HTTP) verifyCallback(c echo.Context) error {
	challenge := c.QueryParam("hub.challenge")
	mode := c.QueryParam("hub.mode")
	verifyToken := c.QueryParam("hub.verify_token")

	// check if the verify token is correct and the mode is subscribe
	if mode == "subscribe" && h.cfg.Server.SubscribeVerifyToken == verifyToken {
		// return the challenge
		return c.String(http.StatusOK, challenge)
	}

	return c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *HTTP) messageEventWebhook(c echo.Context) error {
	waWebhook := &whatsapp.WaWebhook{}

	if err := c.Bind(waWebhook); err != nil {
		return err
	}
	json, err := json.Marshal(waWebhook)
	if err != nil {
		return err
	}

	fmt.Println("bind :- ", string(json))

	return c.JSON(200, http.StatusOK)
}
