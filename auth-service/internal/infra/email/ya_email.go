package email

import (
	"net/smtp"
)

// Prolly won't be used

func ConnectToEmail() {
	from := "moskin.arceniy@yandex.ru" // TODO: все в env
	pass := "akmufvfvxtfmfxan"
	to := []string{"senyamoskin@gmail.com"} // TODO: понять что происходит вообще

	msg := []byte("To: " + to[0] + "\r\n" +
		"Subject: Привет!\r\n" +
		"\r\n" +
		"Hello World! Это тестовое письмо.\r\n")

	addr := "smtp.yandex.ru:587"
	auth := smtp.PlainAuth("", from, pass, "smtp.yandex.ru")

	//c, err := tls.Dial("tcp", addr, nil)
	//
	//if err != nil {
	//	log.Println("Error tcp connecting :", err)
	//	return
	//}
	//
	//client, err := smtp.NewClient(c, "smtp.yandex.ru")
	//
	//if err != nil {
	//	log.Println("Error creating smtp client:", err)
	//	return
	//}
	//
	//err = client.Auth(auth)

	err := smtp.SendMail(addr, auth, from, to, msg) // тут фриз

	if err != nil {
		println(err)
	}

}
