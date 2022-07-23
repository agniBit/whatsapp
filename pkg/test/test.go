package test

import (
	"github.com/agniBit/whatsapp/pkg/whatsapp"
	whatsapp_test "github.com/agniBit/whatsapp/pkg/whatsapp/test"
	"github.com/agniBit/whatsapp/util/config"
)

func Test() error {
	// load config
	cfg, err := config.Load("./util/config/config.yaml")
	if err != nil {
		panic(err)
	}

	// init whatsapp service
	waS := whatsapp.New(cfg)

	// init whatsapp test
	err = whatsapp_test.Test(waS, cfg)

	return err
}
