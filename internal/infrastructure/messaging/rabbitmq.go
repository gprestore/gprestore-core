package messaging

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

func DialRabbitMQ() (*amqp091.Connection, error) {
	conn, err := amqp091.Dial(viper.GetString("rabbitmq.url"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
