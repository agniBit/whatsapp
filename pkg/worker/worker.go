package worker

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Start() error {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092,localhost:9092",
		"group.id":          "worker1",
		"auto.offset.reset": "smallest"},
	)
	if err != nil {
		return err
	}
	run := true

	for run {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			// application-specific processing
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	return nil
}

func SubscribeTopics(consumer *kafka.Consumer, topics []string) error {
	err := consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	run := true

	for run {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(e.Value))
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			fmt.Printf("Ignored %v\n", e)
		}
	}

	consumer.Close()

	return err
}
