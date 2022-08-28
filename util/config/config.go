package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server   *Server   `json:"server" yaml:"server"`
		Whatsapp *Whatsapp `json:"whatsapp" yaml:"whatsapp"`
		Kafka    *Kafka    `json:"kafka" yaml:"kafka"`
	}

	Server struct {
		SubscribeVerifyToken string `json:"subscribe_verify_token" yaml:"subscribe_verify_token"`
	}

	Whatsapp struct {
		Token           string `json:"token" yaml:"token"`
		URL             string `json:"url" yaml:"url"`
		TestPhoneNumber string `json:"test_phone_number" yaml:"test_phone_number"`
	}

	Kafka struct {
		Bootstrap struct {
			Servers string `json:"servers" yaml:"servers"`
		} `json:"bootstrap" yaml:"bootstrap"`
		Security struct {
			Protocol string `json:"protocol" yaml:"security"`
		} `json:"security" yaml:"security"`
		SASL struct {
			Mechanisms string `json:"mechanisms" yaml:"mechanisms"`
			Username   string `json:"username" yaml:"username"`
			Password   string `json:"password" yaml:"password"`
		} `json:"sasl" yaml:"sasl"`
		Acks string `json:"acks" yaml:"acks"`
	}
)

func Load(path string) (*Config, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening the config yaml file %s", err)
	}
	config := &Config{}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return nil, fmt.Errorf("error parsing the config yaml file %v", err)
	}

	//Give first preference to the env fields
	svt := os.Getenv("SUBSCRIBE_VERIFY_TOKEN")
	if svt != "" {
		config.Server.SubscribeVerifyToken = svt
	}

	wt := os.Getenv("WHATSAPP_TOKEN")
	if wt != "" {
		config.Whatsapp.Token = wt
	}

	wtURL := os.Getenv("WHATSAPP_BASE_URL")
	if wtURL != "" {
		config.Whatsapp.URL = wtURL
	}

	wtTestPhoneNumber := os.Getenv("WA_TEST_PHONE")
	if wtTestPhoneNumber != "" {
		config.Whatsapp.TestPhoneNumber = wtTestPhoneNumber
	}

	return config, nil
}

func ReadKafkaConfig(kafkaConfig *Kafka) kafka.ConfigMap {
	m := make(map[string]kafka.ConfigValue)
	m["bootstrap.servers"] = kafkaConfig.Bootstrap.Servers
	m["security.protocol"] = kafkaConfig.Security.Protocol
	m["sasl.mechanisms"] = kafkaConfig.SASL.Mechanisms
	m["sasl.username"] = kafkaConfig.SASL.Username
	m["sasl.password"] = kafkaConfig.SASL.Password
	m["acks"] = kafkaConfig.Acks
	return m
}
