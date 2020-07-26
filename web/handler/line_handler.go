package handler

import (
	"net/http"

	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"
	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotHandler is callback handler for line
type LineBotHandler struct {
	bot     domain.LineBotClient
	service domain.LineService
}

// BuildLineBotHandler creates LineBotHandler struct
func BuildLineBotHandler(container *ioc.Container) *LineBotHandler {
	bot := container.Get((*domain.LineBotClient)(nil)).(domain.LineBotClient)
	lineService := container.Get((*domain.LineService)(nil)).(domain.LineService)

	return &LineBotHandler{bot, lineService}
}

func (h LineBotHandler) handleError(w http.ResponseWriter, err error) {
	switch err {
	case linebot.ErrInvalidSignature:
		w.WriteHeader(400)
	default:
		w.WriteHeader(500)
	}
}

// Index is only used to check sanity
func (h LineBotHandler) Index() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome, here is line bot related path"))
	})
}

// Callback handles line request
func (h LineBotHandler) Callback() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, err := h.bot.ParseRequest(r)
		if err != nil {
			h.handleError(w, err)
			return
		}

		for _, event := range events {
			ctx := &domain.LineContext{Event: event}
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					textMessageContent := &domain.LineTextMessageContext{LineContext: ctx, Message: message}
					h.service.HandleIncomingMessage(textMessageContent)
				}
			}
		}

		w.WriteHeader(200)
	})
}
