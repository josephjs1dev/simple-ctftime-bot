package line

import (
	"net/http"

	"github.com/josephsalimin/simple-ctftime-bot/internal/config"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Bot is our bot implementation that contains actual line-bot-sdk-go
type Bot struct {
	client *linebot.Client
}

// ParseRequest will call line-bot-sdk-go client's ParseRequest
func (bot *Bot) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return bot.client.ParseRequest(r)
}

// ReplyMessage will call line-bot-sdk-go client's ReplyMessage
func (bot *Bot) ReplyMessage(replyToken string, messages ...linebot.SendingMessage) domain.LineBotPushMessageCall {
	return bot.client.ReplyMessage(replyToken, messages...)
}

// InitializeBot initiate line-bot-sdk-go client
func InitializeBot(container *ioc.Container) (domain.LineBotClient, error) {
	config := container.Get((*config.Config)(nil)).(*config.Config)
	client, err := linebot.New(
		config.ChannelSecret,
		config.ChannelToken,
	)

	if err != nil {
		return nil, err
	}

	return &Bot{client: client}, nil
}
