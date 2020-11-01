package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/api/ctftime"
	"github.com/line/line-bot-sdk-go/linebot"
)

// CurrentEventsCommand processes command for current event
type CurrentEventsCommand struct {
	BaseCommand
	client *ctftime.Client
}

// Process ...
func (c *CurrentEventsCommand) Process() ([]linebot.SendingMessage, error) {
	currentEvents, err := c.client.GetCurrentEvents()
	if err != nil && err != ctftime.ErrNoCurrentEvent {
		return nil, err
	}

	if len(currentEvents) == 0 {
		messages := []linebot.SendingMessage{
			linebot.NewTextMessage("No current event"),
		}

		return messages, nil
	}

	contents := make([]*linebot.BubbleContainer, 0)
	// Iterate each object and create carousel template
	for _, currentEvent := range currentEvents {
		contents = append(contents, buildEventBubbleContainer(&currentEvent))
	}

	messages := []linebot.SendingMessage{
		linebot.NewFlexMessage(
			"Current Events Information",
			&linebot.CarouselContainer{
				Type:     linebot.FlexContainerTypeCarousel,
				Contents: contents,
			},
		),
	}

	return messages, nil
}

func buildCurrentEventsCommand(parameter []string) domain.LineCommand {
	return &CurrentEventsCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
		client: ctftime.BuildDefaultClient(),
	}
}
