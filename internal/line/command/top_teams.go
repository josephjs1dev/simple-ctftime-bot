package linecmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/clientapi/ctftime"
	"github.com/line/line-bot-sdk-go/linebot"
	flag "github.com/spf13/pflag"
)

type topTeamsFlag struct {
	year    int
	country string
	total   int
}

// TopTeamsCommand processes command to get top teams information
type TopTeamsCommand struct {
	BaseCommand
	fs         *flag.FlagSet
	flagParams topTeamsFlag
	client     *ctftime.Client
}

// Process ...
func (c *TopTeamsCommand) Process() ([]linebot.SendingMessage, error) {
	topTeams, err := c.client.GetTopTeams(c.flagParams.year, strings.ToUpper(c.flagParams.country), c.flagParams.total)
	if err != nil {
		return nil, err
	}

	contents := make([]linebot.FlexComponent, 0)
	contents = append(contents, &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   "Top Teams",
		Align:  linebot.FlexComponentAlignTypeCenter,
		Weight: linebot.FlexTextWeightTypeBold,
		Size:   linebot.FlexTextSizeTypeLg,
	})

	for idx, team := range topTeams {
		team.Name = fmt.Sprintf("%v. %v", idx+1, team.Name)
		contents = append(contents, buildTopTeamBoxContainer(&team))
	}

	messages := []linebot.SendingMessage{
		linebot.NewFlexMessage(
			"Top Teams",
			&linebot.BubbleContainer{
				Type: linebot.FlexContainerTypeBubble,
				Body: &linebot.BoxComponent{
					Type:     linebot.FlexComponentTypeBox,
					Layout:   linebot.FlexBoxLayoutTypeVertical,
					Contents: contents,
				},
			},
		),
	}

	return messages, nil
}

func buildTopTeamsCommand(parameter []string) domain.LineCommand {
	cmd := &TopTeamsCommand{
		BaseCommand: BaseCommand{
			Parameter: parameter,
		},
		fs:         flag.NewFlagSet("top teams", flag.ContinueOnError),
		flagParams: topTeamsFlag{total: 5},
		client:     ctftime.BuildDefaultClient(),
	}

	cmd.fs.IntVarP(&cmd.flagParams.year, "year", "y", time.Now().Year(), "top teams in chosen year")
	cmd.fs.StringVarP(&cmd.flagParams.country, "country", "c", "", "top teams in chosen country (empty is worlwide)")
	cmd.fs.Parse(parameter)

	return cmd
}
