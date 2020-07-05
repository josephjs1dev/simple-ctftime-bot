package linecontext

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type Context struct {
	Event *linebot.Event
}

type TextMessageContext struct {
	*Context
	Message *linebot.TextMessage
}
