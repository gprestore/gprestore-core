package mq

import (
	"context"
	"log"

	"github.com/gprestore/gprestore-core/internal/infrastructure/messaging"
)

func ConsumeNotificationEmail() error {
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

	for message := range emailConsumer {
		log.Println(string(message.Body))
	}

	return nil
}
