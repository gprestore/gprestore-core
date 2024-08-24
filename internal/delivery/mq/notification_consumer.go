package mq

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gprestore/gprestore-core/internal/infrastructure/messaging"
	"github.com/gprestore/gprestore-core/internal/model"
)

func (c *Consumer) ConsumeNotificationEmail() error {
	conn, err := messaging.DialRabbitMQ()
	if err != nil {
		return err
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	ctx := context.Background()

	emailConsumer, err := channel.ConsumeWithContext(ctx, "email", "email_consumer", true, false, false, false, nil)
	if err != nil {
		return err
	}

	log.Println("Consumer: Consuming Email...")

	for message := range emailConsumer {
		var mail *model.Mail
		err := json.Unmarshal(message.Body, &mail)
		if err != nil {
			return err
		}

		err = c.mailService.Send(mail)
		if err != nil {
			return err
		}
	}

	return nil
}
