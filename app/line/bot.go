package line

import (
	"net/http"
	"simple-ctftime-bot/app/config"

	"github.com/line/line-bot-sdk-go/linebot"
)

// BotClient mimics line-bot-sdk-go Client
type BotClient interface {
	ParseRequest(r *http.Request) ([]*linebot.Event, error)
	ReplyMessage(replyToken string, messages ...linebot.SendingMessage) BotPushMessageCall
}

// BotPushMessageCall mimics line-bot-sdk-go ReplyMessageCall
type BotPushMessageCall interface {
	Do() (*linebot.BasicResponse, error)
}

// AppBot is our bot implementation that contains actual line-bot-sdk-go
type AppBot struct {
	client *linebot.Client
}

// ParseRequest will call line-bot-sdk-go client's ParseRequest
func (bot AppBot) ParseRequest(r *http.Request) ([]*linebot.Event, error) {
	return bot.client.ParseRequest(r)
}

// ReplyMessage will call line-bot-sdk-go client's ReplyMessage
func (bot AppBot) ReplyMessage(replyToken string, messages ...linebot.SendingMessage) BotPushMessageCall {
	return bot.client.ReplyMessage(replyToken, messages...)
}

// InitializeBot initiate line-bot-sdk-go client
func InitializeBot(config config.Config) (BotClient, error) {
	client, err := linebot.New(
		config.ChannelSecret,
		config.ChannelToken,
	)

	if err != nil {
		return nil, err
	}

	return &AppBot{client: client}, nil
}
