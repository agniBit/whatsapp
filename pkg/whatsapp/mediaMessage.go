package whatsapp

import (
	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS waService) SendMediaMessage(templateName, phone_number, imageID string) (*whatsapp.WaResponse, error) {
	templatePayload := whatsapp.WaMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "image",
		Image: &whatsapp.WaImageMessageData{
			ID: imageID,
		},
	}

	return waS.SendMessage(templatePayload)
}
