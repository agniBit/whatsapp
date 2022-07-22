package whatsapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS waService) SendTemplateMessage(templateName, phone_number, parameters string) (*whatsapp.TemplateMessageResponse, error) {
	method := "POST"

	templatePayload := &whatsapp.TemplateMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "template",
		Template: &whatsapp.TemplateInfo{
			Name:     templateName,
			Language: &whatsapp.Language{Code: "en_US"},
		},
	}

	payloadString, err := json.Marshal(*templatePayload)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(payloadString))

	client := &http.Client{}
	req, err := http.NewRequest(method, waS.cgf.Whatsapp.URL, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+waS.cgf.Whatsapp.Token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	resBody := &whatsapp.TemplateMessageResponse{}
	err = json.Unmarshal(body, &resBody)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resBody, nil
}
