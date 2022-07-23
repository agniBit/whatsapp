package whatsapp

import (
	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS WaService) SendMediaMessage(phone_number string, imageData *whatsapp.WaImageMessageData) (*whatsapp.WaResponse, error) {
	templatePayload := whatsapp.WaMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "image",
		Image:            imageData,
	}

	return waS.SendMessage(templatePayload)
}
