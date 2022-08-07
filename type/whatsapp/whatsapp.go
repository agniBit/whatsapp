package whatsapp

type (
	Service interface {
		SendTemplateMessage(templateName, phone_number string, parameters *map[string]string) (*WaResponse, error)
		SendTextMessage(phone_number, message string) (*WaResponse, error)
		SendMediaMessage(phone_number string, imageData *WaImageMessageData) (*WaResponse, error)
		SendInteractiveButtonMessage(phone_number, message string, buttonTexts []string, buttons []*WaInteractiveMessageActionButton) (*WaResponse, error)
		SendInteractiveListMessage(phone_number, header, message, footer, buttonText string, sections []*WaInteractiveMessageActionSection) (*WaResponse, error)
	}

	WaMessagePayload struct {
		MessagingProduct string                    `json:"messaging_product,omitempty"`
		RecipientType    string                    `json:"recipient_type,omitempty"`
		To               string                    `json:"to,omitempty"`
		Type             string                    `json:"type,omitempty"`
		Template         *WaTemplateData           `json:"template,omitempty"`
		Text             *WaTextMessageData        `json:"text,omitempty"`
		Image            *WaImageMessageData       `json:"image,omitempty"`
		Interactive      *WaInteractiveMessageData `json:"interactive,omitempty"`
	}

	WaTemplateData struct {
		Name       string               `json:"name,omitempty"`
		Language   *Language            `json:"language,omitempty"`
		Components []*TemplateComponent `json:"components,omitempty"`
	}

	Language struct {
		Code string `json:"code,omitempty"`
	}

	TemplateComponent struct {
		Type       string              `json:"type,omitempty"`
		Parameters *TemplateParameters `json:"parameters,omitempty"`
	}

	TemplateParameters struct {
		Type     string            `json:"type,omitempty"`
		Text     string            `json:"text,omitempty"`
		Currency string            `json:"currency,omitempty"`
		DateTime *TemplateDateTime `json:"date_time,omitempty"`
	}

	TemplateDateTime struct {
		FallbackValue string `json:"fallback_value,omitempty"`
		DayOfWeek     int    `json:"day_of_week,omitempty"`
		Year          int    `json:"year,omitempty"`
		Month         int    `json:"month,omitempty"`
		Day_of_month  int    `json:"day_of_month,omitempty"`
		Hour          int    `json:"hour,omitempty"`
		Minute        int    `json:"minute,omitempty"`
		Calendar      string `json:"calendar,omitempty"`
	}

	WaResponse struct {
		MessagingProduct string              `json:"messaging_product,omitempty"`
		Contacts         []*TemplateContacts `json:"contacts,omitempty"`
		Message          []*TemplateMessages `json:"messages,omitempty"`
		Error            *Error              `json:"error,omitempty"`
	}

	TemplateContacts struct {
		Input string `json:"input,omitempty"`
		WaId  string `json:"wa_id,omitempty"`
	}

	TemplateMessages struct {
		Id string `json:"id,omitempty"`
	}

	Error struct {
		Message   string `json:"message,omitempty"`
		Type      string `json:"type,omitempty"`
		Code      int    `json:"code,omitempty"`
		FbTraceId string `json:"fbtrace_id,omitempty"`
	}

	WaTextMessageData struct {
		PreviewUrl bool   `json:"preview_url,omitempty"`
		Body       string `json:"body,omitempty"`
	}

	WaImageMessageData struct {
		ID   string `json:"id,omitempty"`
		Link string `json:"link,omitempty"`
	}

	WaInteractiveMessageData struct {
		Type   string                      `json:"type,omitempty"`
		Header *WaInteractiveMessageHeader `json:"header,omitempty"`
		Body   *WaInteractiveMessageBody   `json:"body,omitempty"`
		Footer *WaInteractiveMessageFooter `json:"footer,omitempty"`
		Action *WaInteractiveMessageAction `json:"action,omitempty"`
	}

	WaInteractiveMessageHeader struct {
		Type string `json:"type,omitempty"`
		Text string `json:"text,omitempty"`
	}

	WaInteractiveMessageBody struct {
		Text string `json:"text,omitempty"`
	}

	WaInteractiveMessageFooter struct {
		Text string `json:"text,omitempty"`
	}

	WaInteractiveMessageAction struct {
		Button   string                               `json:"button,omitempty"`
		Buttons  []*WaInteractiveMessageActionButton  `json:"buttons,omitempty"`
		Sections []*WaInteractiveMessageActionSection `json:"sections,omitempty"`
	}

	WaInteractiveMessageActionSection struct {
		Title string                                  `json:"title,omitempty"`
		Rows  []*WaInteractiveMessageActionSectionRow `json:"rows,omitempty"`
	}

	WaInteractiveMessageActionSectionRow struct {
		ID          string `json:"id,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}

	WaInteractiveMessageActionButton struct {
		Type  string                                 `json:"type,omitempty"`
		Reply *WaInteractiveMessageActionButtonReply `json:"reply,omitempty"`
	}

	WaInteractiveMessageActionButtonReply struct {
		ID    string `json:"id,omitempty"`
		Title string `json:"title,omitempty"`
	}
)
