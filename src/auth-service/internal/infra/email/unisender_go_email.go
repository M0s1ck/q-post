package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const unisenderGoUrl = "https://go1.unisender.ru/ru/transactional/api/v1/email/send.json"
const unisenderGoApiKey = "6ubhc6nuxioesbxbhmwx5fsh6cx37znfejsiyhha"

const tempUniosenderGoDomen = "sandbox-7818206-05c2c3.unigosendbox.com"

var tempFromEmail = fmt.Sprintf("auth@%v", tempUniosenderGoDomen)

// Prolly will be used later for email

type UnisenderGoEmailSender struct {
}

type SendEmailRequest struct {
	Message SendEmailMessageType `json:"message"`
}

type SendEmailMessageType struct {
	Recipients []RecipientType   `json:"recipients"`
	Body       SendEmailBodyType `json:"body"`
	Subject    string            `json:"subject"`
	FromEmail  string            `json:"from_email"`
	FromName   string            `json:"from_name"`
}

type RecipientType struct {
	Email string `json:"email"`
}

type SendEmailBodyType struct {
	Html      string `json:"html"`
	PlainText string `json:"plain_text"`
}

func (sender *UnisenderGoEmailSender) SendEmail(to string) {
	requestBodyStruct := SendEmailRequest{
		Message: SendEmailMessageType{
			Recipients: []RecipientType{{Email: to}},
			Subject:    "Q-post auth email",
			FromEmail:  tempFromEmail,
			FromName:   "Q-post auth service",
			Body: SendEmailBodyType{
				Html:      "<b>Hello, world!</b>",
				PlainText: "This is a test email from q-post auth service",
			},
		},
	}

	reqBody, err := json.Marshal(requestBodyStruct)

	if err != nil {
		fmt.Println(err)
		return
	}

	request, err := http.NewRequest("POST", unisenderGoUrl, bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Println(err)
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", unisenderGoApiKey)

	response, err := http.DefaultClient.Do(request) // Rn it's 403 cause domain should be verified

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	bodyString := string(body)
	fmt.Println(bodyString)
}
