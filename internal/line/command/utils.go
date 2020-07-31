package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/line/line-bot-sdk-go/linebot"
)

func buildEventBubbleContainer(event *domain.CTFTimeEvent) *linebot.BubbleContainer {
	descMap := map[string]string{
		"Format":   event.Format,
		"Date":     event.Date,
		"Duration": event.Duration,
		"Team":     event.Team,
	}

	descContents := make([]linebot.FlexComponent, 0)
	for key := range descMap {
		if descMap[key] == "" {
			continue
		}

		content := &linebot.BoxComponent{
			Type:    linebot.FlexComponentTypeBox,
			Layout:  linebot.FlexBoxLayoutTypeBaseline,
			Spacing: linebot.FlexComponentSpacingTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  key,
					Size:  linebot.FlexTextSizeTypeSm,
					Color: "#aaaaaa",
					Flex:  linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: descMap[key],
					Size: linebot.FlexTextSizeTypeSm,
					Flex: linebot.IntPtr(5),
					Wrap: true,
				},
			},
		}

		descContents = append(descContents, content)
	}

	return &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Margin: linebot.FlexComponentMarginTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   event.Title,
					Align:  linebot.FlexComponentAlignTypeCenter,
					Weight: linebot.FlexTextWeightTypeBold,
					Size:   linebot.FlexTextSizeTypeXl,
				},
				&linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Margin:   linebot.FlexComponentMarginTypeLg,
					Spacing:  linebot.FlexComponentSpacingTypeSm,
					Contents: descContents,
				},
			},
		},
		Footer: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.ButtonComponent{
					Style:  linebot.FlexButtonStyleTypeLink,
					Height: linebot.FlexButtonHeightTypeSm,
					Action: linebot.NewURIAction("Open", event.URL),
				},
			},
		},
	}
}
