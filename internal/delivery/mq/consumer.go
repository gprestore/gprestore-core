package mq

import "github.com/gprestore/gprestore-core/internal/service"

type Consumer struct {
	mailService *service.MailService
}

func NewConsumer(mailService *service.MailService) *Consumer {
	return &Consumer{
		mailService: mailService,
	}
}
