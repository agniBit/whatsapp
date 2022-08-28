package whatsapp

import (
	"github.com/agniBit/whatsapp/util/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type WaService struct {
	cgf *config.Config
	pub *kafka.Producer
}

func New(cfg *config.Config, pub *kafka.Producer) WaService {
	return WaService{cgf: cfg, pub: pub}
}
