package whatsapp_test

import (
	"fmt"

	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (t *WaTestService) TestMediaMessages() error {
	// test simple text message
	resp, err := t.waS.SendMediaMessage(t.cfg.Whatsapp.TestPhoneNumber,
		&whatsapp.WaImageMessageData{Link: "https://www.gstatic.com/webp/gallery/1.jpg"},
	)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
