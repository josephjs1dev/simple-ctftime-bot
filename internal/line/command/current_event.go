package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ctftime"
	"github.com/line/line-bot-sdk-go/linebot"
)

// CurrentEventCommand processes command for current event
type CurrentEventCommand struct {
	BaseCommand
	client *ctftime.Client
}

// Process ...
func (c *CurrentEventCommand) Process() ([]linebot.SendingMessage, error) {
	return []linebot.SendingMessage{}, nil
}

func buildCommandEvent(parameter []string) domain.LineCommand {
	return &CurrentEventCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
		client: ctftime.GetClient(),
	}
}
