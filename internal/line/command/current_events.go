package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ctftime"
	"github.com/line/line-bot-sdk-go/linebot"
)

// CurrentEventsCommand processes command for current event
type CurrentEventsCommand struct {
	BaseCommand
	client *ctftime.Client
}

// Process ...
func (c *CurrentEventsCommand) Process() ([]linebot.SendingMessage, error) {
	return []linebot.SendingMessage{}, nil
}

func buildCurrentEventsCommand(parameter []string) domain.LineCommand {
	return &CurrentEventsCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
		client: ctftime.BuildDefaultClient(),
	}
}
