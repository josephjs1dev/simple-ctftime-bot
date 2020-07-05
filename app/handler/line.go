package handler

import (
	"net/http"
	"simple-ctftime-bot/app/ioc"
	appline "simple-ctftime-bot/app/line"
	linecontext "simple-ctftime-bot/app/line/context"
	lineservice "simple-ctftime-bot/app/line/service"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotHandler is callback handler for line
type LineBotHandler struct {
	bot     appline.BotClient
	service lineservice.Service
}

// BuildLineBotHandler creates LineBotHandler struct
func BuildLineBotHandler(container *ioc.Container) *LineBotHandler {
	bot := container.Get((*appline.BotClient)(nil)).(appline.BotClient)
	lineService := container.Get((*lineservice.Service)(nil)).(lineservice.Service)

	return &LineBotHandler{bot, lineService}
}

// Callback handles line request
func (h LineBotHandler) Callback() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, err := h.bot.ParseRequest(r)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			ctx := &linecontext.Context{Event: event}
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					textMessageContent := linecontext.TextMessageContext{
						Context: ctx,
						Message: message,
					}
					err := h.service.HandleIncomingMessage(textMessageContent)
					if err != nil {
						return
					}
				}
			}
		}
	})
}
