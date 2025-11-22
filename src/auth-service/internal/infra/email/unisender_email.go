package email

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Prolly won't be used

type UniSenderEmailSender struct{}

type SendEmailRequestBody struct {
	ApiKey      string `json:"api_key"`
	Email       string `json:"email"`
	SenderName  string `json:"sender_name"`
	SenderEmail string `json:"sender_email"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	ListId      string `json:"list_id"`
}

var url = "https://api.unisender.com/ru/api/sendEmail"

func (sender *UniSenderEmailSender) Send(to string) {
	body := SendEmailRequestBody{
		ApiKey:      "6oqe89dooqkmgiusth1371qxwjyjwhowszntwmro",
		Email:       to,
		SenderName:  "Q-post auth service",
		SenderEmail: "moskin.arceniy@yandex.ru",
		Subject:     "Hello, auth!",
		Body:        "<title>Hello hello</title> <body>This is a test email</body>",
		ListId:      "1", // wth is this
	}

	reqBody, err := json.Marshal(body)

	if err != nil {
		fmt.Println("Json cant do this:", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		fmt.Println("Email post request err:", err)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	dodyStr := string(bodyBytes)
	fmt.Println(dodyStr)
}
