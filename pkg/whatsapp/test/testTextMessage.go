package whatsapp_test

import "fmt"

func (t *WaTestService) TestTextMessages() error {
	// test simple text message
	resp, err := t.waS.SendTextMessage(t.cfg.Whatsapp.TestPhoneNumber, "Hello World")
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
