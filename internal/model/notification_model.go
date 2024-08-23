package model

import "time"

type Notification struct {
	Title        string             `json:"title,omitempty"`
	ShortContent string             `json:"short_content,omitempty"`
	Content      string             `json:"content,omitempty"`
	Target       NotificationTarget `json:"target,omitempty"`
	CreatedAt    *time.Time         `json:"created_at,omitempty"`
}

type NotificationTarget struct {
	Email          string `json:"email,omitempty"`
	DiscordWebhook string `json:"discord_webhook,omitempty"`
}
