package whatsapp_test

import (
	"fmt"

	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (t *WaTestService) TestInteractiveMessages() error {
	// test interactive button messages using payload
	payload := []*whatsapp.WaInteractiveMessageActionButton{
		{
			Type: "reply",
			Reply: &whatsapp.WaInteractiveMessageActionButtonReply{
				ID:    "button_1",
				Title: "Button 1",
			},
		},
		{
			Type: "reply",
			Reply: &whatsapp.WaInteractiveMessageActionButtonReply{
				ID:    "button_2",
				Title: "Button 2",
			},
		},
		{
			Type: "reply",
			Reply: &whatsapp.WaInteractiveMessageActionButtonReply{
				ID:    "button_3",
				Title: "Button 3",
			},
		},
	}

	resp, err := t.waS.SendInteractiveButtonMessage(t.cfg.Whatsapp.TestPhoneNumber, "Hello World", "testing Interactive button message by payload", []string{}, payload)
	if err != nil {
		return err
	}

	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	// send interactive button messages using list of button text
	resp, err = t.waS.SendInteractiveButtonMessage(t.cfg.Whatsapp.TestPhoneNumber, "Hello World", "testing Interactive button message by list of button text", []string{"Button 1", "Button 2", "Button 3"}, nil)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}

	// send interactive list messages using payload
	listPayload := []*whatsapp.WaInteractiveMessageActionSection{
		{
			Title: "Section 1",
			Rows: []*whatsapp.WaInteractiveMessageActionSectionRow{
				{
					ID:          "sec_1_row_1",
					Title:       "Sec 1 Row 1",
					Description: "Sec 1 Description 1",
				},
				{
					ID:          "sec_1_row_2",
					Title:       "Sec 1 Row 2",
					Description: "Sec 1 Description 2",
				},
			},
		},
		{
			Title: "Section 2",
			Rows: []*whatsapp.WaInteractiveMessageActionSectionRow{
				{
					ID:          "sec_2_row_1",
					Title:       "Sec 2 Row 1",
					Description: "Sec 2 Description 1",
				},
				{
					ID:          "sec_2_row_2",
					Title:       "Sec 2 Row 2",
					Description: "Sec 2 Description 2",
				},
			},
		},
	}

	resp, err = t.waS.SendInteractiveListMessage(t.cfg.Whatsapp.TestPhoneNumber, "header", "message", "footer", "Open List", listPayload)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return fmt.Errorf(resp.Error.Message)
	}
	return nil
}
