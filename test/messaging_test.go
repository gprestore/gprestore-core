package test

import (
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/delivery/mq"
	"github.com/gprestore/gprestore-core/internal/gateway/messaging"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/gprestore/gprestore-core/internal/service"
)

func init() {
	config.Load()
}

func TestPublisher(t *testing.T) {
	err := messaging.PublishNotificationEmail(&model.Mail{
		From:    "testing@safatanc.com",
		To:      []string{"agilistikmal3@gmail.com"},
		Subject: "RabbitMQ Testing",
		Body:    "This is rabbitmq testing for gprestore-core",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Publishing message")
}

func TestConsumer(t *testing.T) {
	s := service.NewMailService()
	consumer := mq.NewConsumer(s)
	err := consumer.ConsumeNotificationEmail()
	if err != nil {
		log.Fatal(err)
	}
}
