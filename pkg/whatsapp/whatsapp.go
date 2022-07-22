package whatsapp

import "github.com/agniBit/whatsapp/pkg/gateway"

type HTTP struct {
	svc *gateway.Service
}

func NewHTTP(svc *gateway.Service) *HTTP {
	return &HTTP{
		svc: svc,
	}
}
