package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ctftime"
	"github.com/line/line-bot-sdk-go/linebot"
)

// UpcomingEventsCommand processes command for current event
type UpcomingEventsCommand struct {
	BaseCommand
	client *ctftime.Client
}

// Process ...
func (c *UpcomingEventsCommand) Process() ([]linebot.SendingMessage, error) {
	upcomingEvents, err := c.client.GetUpcomingEvents()
	if err != nil {
		return nil, err
	}

	contents := make([]*linebot.BubbleContainer, 0)
	// Iterate each object and create carousel template
	for _, upcomingEvent := range upcomingEvents {
		contents = append(contents, buildEventBubbleContainer(&upcomingEvent))
	}

	messages := []linebot.SendingMessage{
		linebot.NewFlexMessage(
			"Upcoming Events Information",
			&linebot.CarouselContainer{
				Type:     linebot.FlexContainerTypeCarousel,
				Contents: contents,
			},
		),
	}

	return messages, nil
}

func buildUpcomingEventsCommand(parameter []string) domain.LineCommand {
	return &UpcomingEventsCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
		client: ctftime.BuildDefaultClient(),
	}
}
