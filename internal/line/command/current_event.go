package linecmd

import "github.com/line/line-bot-sdk-go/linebot"

type CurrentEventCommand struct {
	BaseCommand
}

func (c *CurrentEventCommand) Process() ([]linebot.SendingMessage, error) {
	return []linebot.SendingMessage{}, nil
}

func buildCommandEvent(parameter []string) Command {
	return &CurrentEventCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
	}
}
