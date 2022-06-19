package discord

import (
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/rest"
	"github.com/disgoorg/disgo/webhook"
	"github.com/disgoorg/log"
	"github.com/disgoorg/snowflake/v2"
)

type WebHook struct {
	WebhookID uint64
	Token     string
	UserId    uint64
}

func (w *WebHook) Send(s string) {
	client := webhook.New(snowflake.ID(w.WebhookID), w.Token)

	if _, err := client.CreateMessage(discord.NewWebhookMessageCreateBuilder().
		SetContentf("<@%d> %s", w.UserId, s).
		Build(),
		// delay each request by 1 seconds
		rest.WithDelay(1*time.Second),
	); err != nil {
		log.Errorf("error sending message %s: %s", s, err)
	}
}
