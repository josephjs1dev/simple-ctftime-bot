package linecmd

import "github.com/line/line-bot-sdk-go/linebot"

// Command is interface for line command
type Command interface {
	Process() ([]linebot.SendingMessage, error)
}

// BaseCommand will be embedded by all children command
type BaseCommand struct {
	Parameter []string
}
