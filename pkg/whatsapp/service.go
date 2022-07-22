package whatsapp

import "github.com/agniBit/whatsapp/util/config"

type waService struct {
	cgf *config.Config
}

func New(cfg *config.Config) waService {
	return waService{cgf: cfg}
}
