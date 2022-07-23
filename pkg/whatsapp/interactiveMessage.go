package whatsapp

import (
	"strconv"

	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS WaService) SendInteractiveButtonMessage(templateName, phone_number, message string, buttonTexts []string, buttons []*whatsapp.WaInteractiveMessageActionButton) (*whatsapp.WaResponse, error) {
	if len(buttons) == 0 {
		buttons = []*whatsapp.WaInteractiveMessageActionButton{}

		for index, buttonText := range buttonTexts {
			buttons = append(buttons, &whatsapp.WaInteractiveMessageActionButton{
				Type: "reply",
				Reply: &whatsapp.WaInteractiveMessageActionButtonReply{
					ID:    "button_" + strconv.Itoa(index),
					Title: buttonText,
				},
			})
		}
	}

	templatePayload := whatsapp.WaMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "interactive",
		Interactive: &whatsapp.WaInteractiveMessageData{
			Type: "button",
			Body: &whatsapp.WaInteractiveMessageBody{
				Text: message,
			},
			Action: &whatsapp.WaInteractiveMessageAction{
				Buttons: buttons,
			},
		},
	}

	return waS.SendMessage(templatePayload)
}

func (waS WaService) SendInteractiveListMessage(phone_number, header, message, footer, buttonText string, sections []*whatsapp.WaInteractiveMessageActionSection) (*whatsapp.WaResponse, error) {
	templatePayload := whatsapp.WaMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "interactive",
		Interactive: &whatsapp.WaInteractiveMessageData{
			Type: "list",
			Header: &whatsapp.WaInteractiveMessageHeader{
				Text: header,
			},
			Body: &whatsapp.WaInteractiveMessageBody{
				Text: message,
			},
			Footer: &whatsapp.WaInteractiveMessageFooter{
				Text: footer,
			},
			Action: &whatsapp.WaInteractiveMessageAction{
				Button:   buttonText,
				Sections: sections,
			},
		},
	}

	return waS.SendMessage(templatePayload)
}
