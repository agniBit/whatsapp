package whatsapp

import "github.com/agniBit/whatsapp/model"

type (
	Service interface {
		SendTemplateMessage(templateName, phone_number, parameters string) (*model.TemplateMessageResponse, error)
	}
)
