package whatsapp

import (
	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS waService) SendTemplateMessage(templateName, phone_number, parameters string) (*whatsapp.WaResponse, error) {
	templatePayload := whatsapp.WaMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "template",
		Template: &whatsapp.WaTemplateData{
			Name:     templateName,
			Language: &whatsapp.Language{Code: "en_US"},
		},
	}

	return waS.SendMessage(templatePayload)
}
