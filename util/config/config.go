package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Server   *Server   `json:"server" yaml:"server"`
		Whatsapp *Whatsapp `json:"whatsapp" yaml:"whatsapp"`
	}

	Server struct {
		SubscribeVerifyToken string `json:"subscribe_verify_token" yaml:"subscribe_verify_token"`
	}

	Whatsapp struct {
		Token string `json:"token" yaml:"token"`
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

	return config, nil
}
