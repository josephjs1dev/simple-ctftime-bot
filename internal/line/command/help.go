package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/line/line-bot-sdk-go/linebot"
)

// HelpCommand returns commands information
type HelpCommand struct {
	BaseCommand
}

// Process ...
func (c *HelpCommand) Process() ([]linebot.SendingMessage, error) {
	commandMap := map[string]string{
		"!help":            "to get commands usage",
		"!upcoming_events": "to get upcomings CTFTime events",
	}

	commandContents := make([]linebot.FlexComponent, 0)
	for key := range commandMap {
		content := &linebot.BoxComponent{
			Type:    linebot.FlexComponentTypeBox,
			Layout:  linebot.FlexBoxLayoutTypeVertical,
			Spacing: linebot.FlexComponentSpacingTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  key,
					Size:  linebot.FlexTextSizeTypeSm,
					Color: "#aaaaaa",
					Wrap:  true,
				},
				&linebot.BoxComponent{
					Type:    linebot.FlexComponentTypeBox,
					Layout:  linebot.FlexBoxLayoutTypeBaseline,
					Spacing: linebot.FlexComponentSpacingTypeSm,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: "  -",
							Size: linebot.FlexTextSizeTypeSm,
							Wrap: true,
							Flex: linebot.IntPtr(1),
						},
						&linebot.TextComponent{
							Type: linebot.FlexComponentTypeText,
							Text: commandMap[key],
							Size: linebot.FlexTextSizeTypeSm,
							Wrap: true,
							Flex: linebot.IntPtr(10),
						},
					},
				},
			},
		}

		commandContents = append(commandContents, content)
	}

	contents := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Margin: linebot.FlexComponentMarginTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: "List of available commands",
				},
				&linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Margin:   linebot.FlexComponentMarginTypeLg,
					Spacing:  linebot.FlexComponentSpacingTypeSm,
					Contents: commandContents,
				},
			},
		},
	}

	messages := []linebot.SendingMessage{linebot.NewFlexMessage("Upcoming Events Information", contents)}

	return messages, nil
}

func buildHelpCommand(parameter []string) domain.LineCommand {
	return &HelpCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
	}
}
