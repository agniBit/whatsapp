package whatsapp_test

import "fmt"

func (t *WaTestService) TestTemplateMessages() error {
	// test simple text message
	parameters := &map[string]string{}
	resp, err := t.waS.SendTemplateMessage("hello_world", t.cfg.Whatsapp.TestPhoneNumber, parameters)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	return nil
}
