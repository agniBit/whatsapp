package event

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"

	"github.com/agniBit/whatsapp/pkg/worker"
	"github.com/agniBit/whatsapp/type/whatsapp"
	"github.com/agniBit/whatsapp/util/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type event struct {
	cgf *config.Config
}

func NewKafkaEvent(cgf *config.Config) event {
	return event{
		cgf: cgf,
	}
}

func (ke event) Consumer() error {
	conf := config.ReadKafkaConfig(ke.cgf.Kafka)

	topic := string(worker.KafkaTopicWhatsAppMessage)
	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		return err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Set up a channel for handling Ctrl-C, etc
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// Process messages
	run := true

	// check terminate signal
	go func() {
		<-quit
		run = false
	}()

	for run {
		ev := c.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			ke.handlerKafkaMessageEvent(e)
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	c.Close()

	return nil
}

func (ke event) handlerKafkaMessageEvent(e *kafka.Message) {
	fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
		*e.TopicPartition.Topic, string(e.Key), string(e.Value))

	payload := &whatsapp.WaMessagePayload{}
	err := json.Unmarshal(e.Value, payload)
	if err != nil {
		return
	}

	_, err = ke.sendWhatsappMessage(*payload)

	if err != nil {
		fmt.Println(err.Error())
	}
}
