package whatsapp

type (
	WaWebhook struct {
		Object string     `json:"object,omitempty"`
		Entry  []*WaEntry `json:"entry,omitempty"`
	}

	WaEntry struct {
		ID      string           `json:"id,omitempty"`
		Changes []*WaEntryChange `json:"changes,omitempty"`
	}

	WaEntryChange struct {
		Value *WaEntryChangeValue `json:"value,omitempty"`
		Field string              `json:"field,omitempty"`
	}

	WaEntryChangeValue struct {
		MessagingProduct string       `json:"messaging_product,omitempty"`
		MetaData         WaMetaData   `json:"metadata,omitempty"`
		Contacts         []*WaContact `json:"contacts,omitempty"`
		Message          []*WaMessage `json:"messages,omitempty"`
	}

	WaMetaData struct {
		DisplayPhoneNumber string `json:"display_phone_number,omitempty"`
		PhoneNumberID      string `json:"phone_number_id,omitempty"`
	}

	WaContact struct {
		Profile struct {
			Name string `json:"name,omitempty"`
		} `json:"profile,omitempty"`
		WaID string `json:"Wa_id,omitempty"`
	}

	WaMessage struct {
		ID          string                `json:"id,omitempty"`
		From        string                `json:"from,omitempty"`
		Timestamp   string                `json:"timestamp,omitempty"`
		Type        string                `json:"type,omitempty"`
		Text        *WaText               `json:"text,omitempty"`
		Interactive *WaInteractiveMessage `json:"interactive,omitempty"`
		Image       *WaImageMessage       `json:"image,omitempty"`
		Sticker     *WaStickerMessage     `json:"sticker,omitempty"`
		Location    *WaLocationMessage    `json:"location,omitempty"`
		Button      *WaButtonMessage      `json:"button,omitempty"`
		Context     *WaMessageContext     `json:"context,omitempty"`
	}

	WaText struct {
		Body string `json:"body,omitempty"`
	}

	WaMessageContext struct {
		ID   string `json:"id,omitempty"`
		From string `json:"from,omitempty"`
	}

	WaInteractiveMessage struct {
		Type        string         `json:"type,omitempty"`
		ButtonReply *WaButtonReply `json:"button_reply,omitempty"`
		ListReply   *WaListReply   `json:"list_reply,omitempty"`
	}

	WaButtonReply struct {
		ID    string `json:"id,omitempty"`
		Title string `json:"title,omitempty"`
	}

	WaListReply struct {
		ID          string `json:"id,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}

	WaImageMessage struct {
		ID       string `json:"id,omitempty"`
		Caption  string `json:"caption,omitempty"`
		MimeType string `json:"mime_type,omitempty"`
		SHA256   string `json:"sha256,omitempty"`
	}

	WaStickerMessage struct {
		ID       string `json:"id,omitempty"`
		MimeType string `json:"mime_type,omitempty"`
		SHA256   string `json:"sha256,omitempty"`
	}

	WaLocationMessage struct {
		Name    string  `json:"name,omitempty"`
		Address string  `json:"address,omitempty"`
		Lat     float64 `json:"latitude,omitempty"`
		Lon     float64 `json:"longitude,omitempty"`
	}

	WaButtonMessage struct {
		Text    string `json:"text,omitempty"`
		Payload string `json:"payload,omitempty"`
	}
)
