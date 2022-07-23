package whatsapp_test

import (
	"github.com/agniBit/whatsapp/pkg/whatsapp"
	"github.com/agniBit/whatsapp/util/config"
)

type WaTestService struct {
	cfg *config.Config
	waS whatsapp.WaService
}

func TestWhatsapp(waS whatsapp.WaService, cfg *config.Config) error {
	waTestService := &WaTestService{waS: waS, cfg: cfg}

	// test template messages
	err := waTestService.TestTemplateMessages()
	if err != nil {
		return err
	}

	// test simple text messages
	err = waTestService.TestTextMessages()
	if err != nil {
		return err
	}

	// test interactive messages
	err = waTestService.TestInteractiveMessages()
	if err != nil {
		return err
	}

	// test media messages
	err = waTestService.TestMediaMessages()
	if err != nil {
		return err
	}

	return nil
}
