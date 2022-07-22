package gateway

import (
	"github.com/agniBit/whatsapp/type/whatsapp"
)

type Service struct {
	waS whatsapp.Service
}

func New(
	waS whatsapp.Service,
) Service {
	return Service{
		waS: waS,
	}
}
