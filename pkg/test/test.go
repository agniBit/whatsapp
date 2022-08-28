package test

import (
	"github.com/agniBit/whatsapp/pkg/whatsapp"
	whatsapp_test "github.com/agniBit/whatsapp/pkg/whatsapp/test"
	"github.com/agniBit/whatsapp/util/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Test() error {
	// load config
	cfg, err := config.Load("./util/config/config.yaml")
	if err != nil {
		panic(err)
	}

	// create new kafka publisher
	kafkaConfig := config.ReadKafkaConfig(cfg.Kafka)
	pub, err := kafka.NewProducer(&kafkaConfig)
	if err != nil {
		panic(err)
	}

	// init whatsapp service
	waS := whatsapp.New(cfg, pub)

	// init whatsapp test
	err = whatsapp_test.Test(waS, cfg)

	return err
}
