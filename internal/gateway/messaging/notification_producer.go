package messaging

import (
	"context"
	"encoding/json"

	"github.com/gprestore/gprestore-core/internal/infrastructure/messaging"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/rabbitmq/amqp091-go"
)

func PublishNotificationEmail(mail *model.Mail) error {
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

	mailJson, err := json.Marshal(mail)
	if err != nil {
		return err
	}

	message := amqp091.Publishing{
		Body: mailJson,
	}

	return channel.PublishWithContext(ctx, "notification", "email", false, false, message)
}
