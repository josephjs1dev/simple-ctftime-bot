package domain

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
)

// LineBotClient mimics line-bot-sdk-go Client
type LineBotClient interface {
	ParseRequest(*http.Request) ([]*linebot.Event, error)
	ReplyMessage(string, ...linebot.SendingMessage) LineBotPushMessageCall
}

// LineBotPushMessageCall mimics line-bot-sdk-go ReplyMessageCall
type LineBotPushMessageCall interface {
	Do() (*linebot.BasicResponse, error)
}

// LineContext is base context that has Event type properties
type LineContext struct {
	*linebot.Event
}

// LineTextMessageContext embedded Context and add TextMessage type properties
type LineTextMessageContext struct {
	*LineContext
	Message *linebot.TextMessage
}

// LineService is our line service interface that defines function that needs to be implemented
type LineService interface {
	HandleIncomingMessage(*LineTextMessageContext) error
}

// LineCommand is interface for line command
type LineCommand interface {
	Process() ([]linebot.SendingMessage, error)
}

// LineCommandBuilder is builder function to generate command processor
type LineCommandBuilder func([]string) LineCommand

// LineCommandMapper is mapper that returns LineCommandBuilder
type LineCommandMapper interface {
	GetCommand(string) LineCommandBuilder
}
