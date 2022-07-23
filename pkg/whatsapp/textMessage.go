package whatsapp

import (
	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS waService) SendTextMessage(templateName, phone_number, message string) (*whatsapp.WaResponse, error) {
	textPayload := whatsapp.WaMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "text",
		Text: &whatsapp.WaTextMessageData{
			PreviewUrl: true,
			Body:       message,
		},
	}

	return waS.SendMessage(textPayload)
}
