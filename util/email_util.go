package util

import (
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
)

func EmailUtil() {
	// ch := make(chan *email.Email, 10)
	p, err := email.NewPool(
		"smtp.163.com:465",
		4,
		smtp.PlainAuth("", "new001code@163.com", "PTFKEIEKJFGFODJI", "smtp.163.com"),
	)
	if err != nil {
		ErrorLogger.Fatalln("email init error")
	}

	e := email.NewEmail()
	e.From = "new001cide@163.com"
	e.To = []string{"lhl_creeper@163.com"}
	e.Subject = "test"
	e.Text = []byte("test context")
	p.Send(e, 10*time.Second)
}
