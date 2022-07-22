package whatsapp

type (
	Service interface {
		SendTemplateMessage(templateName, phone_number, parameters string) (*TemplateMessageResponse, error)
	}
	TemplateMessagePayload struct {
		MessagingProduct string        `json:"messaging_product"`
		RecipientType    string        `json:"recipient_type"`
		To               string        `json:"to"`
		Type             string        `json:"type"`
		Template         *TemplateInfo `json:"template"`
	}

	TemplateInfo struct {
		Name       string               `json:"name"`
		Language   *Language            `json:"language"`
		Components []*TemplateComponent `json:"components,omitempty"`
	}

	Language struct {
		Code string `json:"code"`
	}

	TemplateComponent struct {
		Type       string              `json:"type,omitempty"`
		Parameters *TemplateParameters `json:"parameters,omitempty"`
	}

	TemplateParameters struct {
		Type     string             `json:"type"`
		Text     string             `json:"text,omitempty"`
		Currency string             `json:"currency,omitempty"`
		DateTime *TeamplateDateTime `json:"date_time,omitempty"`
	}

	TeamplateDateTime struct {
		FallbackValue string `json:"fallback_value,omitempty"`
		Dayofweek     int    `json:"day_of_week,omitempty"`
		Year          int    `json:"year,omitempty"`
		Month         int    `json:"month,omitempty"`
		Day_of_month  int    `json:"day_of_month,omitempty"`
		Hour          int    `json:"hour,omitempty"`
		Minute        int    `json:"minute,omitempty"`
		Calendar      string `json:"calendar,omitempty"`
	}

	TemplateMessageResponse struct {
		MessagingProduct string              `json:"messaging_product"`
		Contacts         []*TemplateContacts `json:"contacts"`
		Message          []*TemplateMessages `json:"messages"`
		Error            *Error              `json:"error,omitempty"`
	}

	TemplateContacts struct {
		Input string `json:"input"`
		WaId  string `json:"wa_id"`
	}

	TemplateMessages struct {
		Id string `json:"id"`
	}

	Error struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceId string `json:"fbtrace_id"`
	}
)
