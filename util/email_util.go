package util

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/jordan-wright/email"
	"github.com/spf13/viper"

	"ginchat/common"
)

type EmailUtil struct{}

var (
	ch       = make(chan *email.Email, 8)
	pool     *email.Pool
	fromName = fmt.Sprintf("%s <%s>", viper.GetString("mail.fromName"), viper.GetString("mail.username"))
)

func init() {
	common.Logger.Info("email pool init")
	h := fmt.Sprintf("%s:%d", viper.GetString("mail.host"), viper.GetInt("mail.port"))
	ps := viper.GetInt("mail.poolSize")
	un := viper.GetString("mail.username")
	pw := viper.GetString("mail.password")
	common.Logger.Infof("host:%s", h)
	mailPool, err := email.NewPool(
		h,
		ps,
		smtp.PlainAuth("", un, pw, viper.GetString("mail.host")),
	)
	if err != nil {
		common.Logger.Error("fail create mail pool", err)
	}
	pool = mailPool
}

func sendEmail() {
	e := <-ch
	if err := pool.Send(e, 10*time.Second); err != nil {
		common.Logger.Error("send email err: to: %s  err: %s", e.To, err)
	} else {
		common.Logger.Debug("send email success")
	}
}

func (mu *EmailUtil) SendText(receiver []string, subject string, text string) {
	e := email.NewEmail()
	e.From = fromName
	e.To = receiver
	e.Subject = subject
	e.Text = []byte(text)
	ch <- e
	go sendEmail()
}

func (mu *EmailUtil) SendHTML(receiver []string, subject string, html string) {
	e := email.NewEmail()
	e.From = fromName
	e.To = receiver
	e.Subject = subject
	e.HTML = []byte(html)
	ch <- e
	go sendEmail()
}
