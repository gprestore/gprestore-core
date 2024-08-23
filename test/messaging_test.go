package test

import (
	"log"
	"testing"

	"github.com/gprestore/gprestore-core/internal/config"
	"github.com/gprestore/gprestore-core/internal/delivery/mq"
	"github.com/gprestore/gprestore-core/internal/gateway/messaging"
	"github.com/gprestore/gprestore-core/internal/model"
)

func init() {
	config.Load()
}

func TestPublisher(t *testing.T) {
	err := messaging.PublishNotification(&model.Notification{
		Title:        "Judul",
		ShortContent: "Deskripsi Singkat",
		Content:      "Deskripsi Panjang",
		Target: model.NotificationTarget{
			Email: "agilistikmal3@gmail.com",
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func TestConsumer(t *testing.T) {
	err := mq.ConsumeNotificationEmail()
	if err != nil {
		log.Fatal(err)
	}
}
