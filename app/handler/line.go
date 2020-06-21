package handler

import (
	"net/http"
	"simple-ctftime-bot/app/content"
	"simple-ctftime-bot/app/line"
	linecontent "simple-ctftime-bot/app/line/content"
	lineservice "simple-ctftime-bot/app/line/service"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotHandler is callback handler for line
type LineBotHandler struct {
	bot     line.BotClient
	service lineservice.Service
}

// BuildLineBotHandler creates LineBotHandler struct
func BuildLineBotHandler(content *content.AppContent, service lineservice.Service) *LineBotHandler {
	return &LineBotHandler{bot: content.Line, service: service}
}

// Callback handles line request
func (h LineBotHandler) Callback(w http.ResponseWriter, r *http.Request) {
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
		lineContent := &linecontent.Content{Event: event}
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				textMessageContent := linecontent.TextMessageContent{
					Content: lineContent,
					Message: message,
				}
				err := h.service.HandleIncomingMessage(textMessageContent)
				if err != nil {
					return
				}
			}
		}
	}
}
