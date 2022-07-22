package whatsapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/agniBit/whatsapp/model"
)

func (waS waService) SendTemplateMessage(templateName, phone_number, parameters string) (*model.TemplateMessageResponse, error) {
	url := "https://graph.facebook.com/v13.0/104569905665986/messages"
	method := "POST"

	templatePayload := &model.TemplateMessagePayload{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               phone_number,
		Type:             "template",
		Template: &model.TemplateInfo{
			Name:     templateName,
			Language: &model.Language{Code: "en_US"},
		},
	}

	payloadString, err := json.Marshal(*templatePayload)
	if err != nil {
		return nil, err
	}

	payload := strings.NewReader(string(payloadString))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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

	resBody := &model.TemplateMessageResponse{}
	err = json.Unmarshal(body, &resBody)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resBody, nil
}
