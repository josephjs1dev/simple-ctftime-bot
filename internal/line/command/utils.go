package linecmd

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/line/line-bot-sdk-go/linebot"
)

type flexDescComponent struct {
	key     string
	value   string
	keySize linebot.FlexTextSizeType
}

func getFlexTextSize(size linebot.FlexTextSizeType, defaultValue linebot.FlexTextSizeType) linebot.FlexTextSizeType {
	if size == "" {
		return defaultValue
	}

	return size
}

func buildEventBubbleContainer(event *domain.CTFTimeEvent) *linebot.BubbleContainer {
	descComponents := []flexDescComponent{
		{key: "Format", value: event.Format},
		{key: "Date", value: event.Date},
		{key: "Duration", value: event.Duration},
		{key: "Team", value: event.Team},
	}

	descContents := make([]linebot.FlexComponent, 0)
	for _, component := range descComponents {
		if component.value == "" {
			continue
		}

		content := &linebot.BoxComponent{
			Type:    linebot.FlexComponentTypeBox,
			Layout:  linebot.FlexBoxLayoutTypeHorizontal,
			Spacing: linebot.FlexComponentSpacingTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  component.key,
					Size:  linebot.FlexTextSizeTypeSm,
					Color: "#aaaaaa",
					Flex:  linebot.IntPtr(2),
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: component.value,
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

func buildTopTeamBoxContainer(team *domain.CTFTimeTeam) linebot.FlexComponent {
	descComponents := []flexDescComponent{
		{key: "Worldwide Position", value: team.WorldwidePosition},
		{key: "Points", value: team.Points},
		{key: "Events", value: team.Events},
	}

	descContents := make([]linebot.FlexComponent, 0)
	for _, component := range descComponents {
		if component.value == "" {
			continue
		}

		content := &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeHorizontal,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: " ",
					Flex: linebot.IntPtr(1),
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  component.key,
					Size:  linebot.FlexTextSizeTypeXs,
					Color: "#aaaaaa",
					Flex:  linebot.IntPtr(10),
					Wrap:  true,
				},
				&linebot.TextComponent{
					Type: linebot.FlexComponentTypeText,
					Text: component.value,
					Size: linebot.FlexTextSizeTypeXs,
					Flex: linebot.IntPtr(8),
					Wrap: true,
				},
			},
		}

		descContents = append(descContents, content)
	}

	return &linebot.BoxComponent{
		Type:    linebot.FlexComponentTypeBox,
		Layout:  linebot.FlexBoxLayoutTypeVertical,
		Margin:  linebot.FlexComponentMarginTypeXxl,
		Spacing: linebot.FlexComponentSpacingTypeNone,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   team.Name,
				Weight: linebot.FlexTextWeightTypeBold,
				Size:   linebot.FlexTextSizeTypeMd,
			},
			&linebot.BoxComponent{
				Type:     linebot.FlexComponentTypeBox,
				Layout:   linebot.FlexBoxLayoutTypeVertical,
				Margin:   linebot.FlexComponentMarginTypeSm,
				Spacing:  linebot.FlexComponentSpacingTypeNone,
				Contents: descContents,
			},
		},
	}
}
