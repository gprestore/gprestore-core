package messaging

import (
	"context"
	"encoding/json"

	"github.com/gprestore/gprestore-core/internal/infrastructure/messaging"
	"github.com/gprestore/gprestore-core/internal/model"
	"github.com/rabbitmq/amqp091-go"
)

func PublishNotification(notification *model.Notification) error {
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

	notificationJson, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	message := amqp091.Publishing{
		Body: notificationJson,
	}

	return channel.PublishWithContext(ctx, "notification", "email", false, false, message)
}
