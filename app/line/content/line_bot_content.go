package linecontent

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

type Content struct {
	Event *linebot.Event
}

type TextMessageContent struct {
	*Content
	Message *linebot.TextMessage
}
