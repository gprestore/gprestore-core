package service

import (
	"crypto/tls"
	"fmt"

	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type MailService struct {
}

func NewMailService() *MailService {
	return &MailService{}
}

func (s *MailService) Dial(from string) (*gomail.Dialer, error) {
	accounts := viper.Get("mail.accounts")

	var username string
	var password string
	for _, account := range accounts.([]any) {
		u := account.(map[string]any)["username"].(string)
		p := account.(map[string]any)["password"].(string)

		if u == from {
			username = u
			password = p
			break
		}
	}

	if username == "" && password == "" {
		return nil, fmt.Errorf("config mail password for %s not found", from)
	}

	conn := gomail.NewDialer(
		viper.GetString("mail.smtp.host"),
		viper.GetInt("mail.smtp.port"),
		username,
		password,
	)
	conn.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return conn, nil
}

func (s *MailService) Send(mail *model.Mail) error {
	m := gomail.NewMessage()
	m.SetHeader("From", mail.From)
	m.SetHeader("To", mail.To...)

	if mail.Cc != nil {
		m.SetAddressHeader("Cc", mail.Cc.Address, mail.Cc.Name)
	}

	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	conn, err := s.Dial(mail.From)
	if err != nil {
		return err
	}

	return conn.DialAndSend(m)
}
