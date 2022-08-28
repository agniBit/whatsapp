package whatsapp

import (
	"encoding/json"

	"github.com/agniBit/whatsapp/pkg/worker"
	"github.com/agniBit/whatsapp/type/whatsapp"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (waS WaService) SendMessage(payload whatsapp.WaMessagePayload) (*whatsapp.WaResponse, error) {
	val, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	topic := string(worker.KafkaTopicWhatsAppMessage)

	err = waS.pub.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          val,
	}, nil)

	return nil, err
}
