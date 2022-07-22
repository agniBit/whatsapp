package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/agniBit/whatsapp/pkg/gateway"
	httpS "github.com/agniBit/whatsapp/pkg/gateway/transport"
	"github.com/agniBit/whatsapp/pkg/whatsapp"
	"github.com/labstack/echo/v4"
)

func Start() error {
	e := echo.New()
	v1 := e.Group("/api/v1")

	// init whatsapp service
	waS := whatsapp.New()

	// init gateway service
	gatewayS := gateway.New(waS)

	// init gateway transport
	httpS.NewHTTP(gatewayS, v1)

	// start server
	StartServer(e)

	return nil
}

func StartServer(e *echo.Echo) {
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", 8080),
	}

	// Start server
	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Info("Shutting down the server", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
