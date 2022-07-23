package whatsapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/agniBit/whatsapp/type/whatsapp"
)

func (waS waService) SendMessage(payload whatsapp.WaMessagePayload) (*whatsapp.WaResponse, error) {
	method := "POST"

	payloadString, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	payloadReader := strings.NewReader(string(payloadString))

	client := &http.Client{}
	req, err := http.NewRequest(method, waS.cgf.Whatsapp.URL, payloadReader)

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

	resBody := &whatsapp.WaResponse{}
	err = json.Unmarshal(body, &resBody)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resBody, nil
}
