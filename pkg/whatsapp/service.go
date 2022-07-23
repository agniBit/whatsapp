package whatsapp

import "github.com/agniBit/whatsapp/util/config"

type WaService struct {
	cgf *config.Config
}

func New(cfg *config.Config) WaService {
	return WaService{cgf: cfg}
}
