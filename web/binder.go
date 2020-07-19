package web

import (
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
	"github.com/josephsalimin/simple-ctftime-bot/internal/line"
	linecmd "github.com/josephsalimin/simple-ctftime-bot/internal/line/command"
	lineservice "github.com/josephsalimin/simple-ctftime-bot/internal/line/service"
	"github.com/josephsalimin/simple-ctftime-bot/internal/pkg/ioc"
)

type implementationBinder func(*ioc.Container) error

func bindLineBot(c *ioc.Container) error {
	client, err := line.InitializeBot(c)
	if err != nil {
		return err
	}

	return c.BindInterface(client, (*domain.LineBotClient)(nil))
}

func bindLineCmd(c *ioc.Container) error {
	lineCmdMapper := linecmd.BuildCommandMapper()
	return c.BindInterface(lineCmdMapper, (*domain.LineCommandMapper)(nil))
}

func bindLineService(c *ioc.Container) error {
	lineService := lineservice.BuildService(c)
	return c.BindInterface(lineService, (*domain.LineService)(nil))
}

var binders []implementationBinder = []implementationBinder{
	bindLineBot,
	bindLineCmd,
	bindLineService,
}
