package lineservice

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// Context is base context that has Event type properties
type Context struct {
	Event *linebot.Event
}

// TextMessageContext embedded Context and add TextMessage type properties
type TextMessageContext struct {
	*Context
	Message *linebot.TextMessage
}
