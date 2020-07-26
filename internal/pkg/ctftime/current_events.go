package ctftime

import (
	"github.com/anaskhan96/soup"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
)

var currentEventTextOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"div", "class", "span6"}, findIndex: 1},
	{findType: findOne, findParams: []string{"div", "class", "page-header"}},
	{findType: findOneInAll, findParams: []string{"a"}, findIndex: 0},
}

var currentEventsTraversalOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"div", "class", "span6"}, findIndex: 1},
	{findType: findOne, findParams: []string{"table", "width", "100%"}},
	{findType: findOne, findParams: []string{"tbody"}},
	{findType: findAll, findParams: []string{"tr"}},
}

var currentEventFormatOpts = []htmlTraversalOption{
	{findType: findOne, findParams: []string{"img"}},
}

var currentEventTitleOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"a"}, findIndex: 1},
}

var currentEventDateOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 1},
}

var currentEventDurationOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 1},
	{findType: findOne, findParams: []string{"small", "class", "muted pull-right"}},
}

var currentEventTeamOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 1},
}

func checkCurrentEventText(node soup.Root) (bool, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventTextOpts)
	if err != nil {
		return false, err
	}

	return child[0].Text() == "Current Events", nil
}

func getCurrentEventFormat(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventFormatOpts)
	if err != nil {
		return "", err
	}

	return getAttrKey(child[0], "title"), nil
}

func getCurrentEventTitle(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventTitleOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getCurrentEventURI(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventTitleOpts)
	if err != nil {
		return "", err
	}

	return getAttrKey(child[0], "href"), nil
}

func getCurrentEventDate(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventDateOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getCurrentEventDuration(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventDateOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getCurrentEventTeam(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, currentEventTeamOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

// GetCurrentEvents ...
func (c *Client) GetCurrentEvents() ([]domain.CTFTimeEvent, error) {
	upcomingEvents := make([]domain.CTFTimeEvent, 0)

	body, err := c.Get(c.baseURL)
	if err != nil {
		return nil, err
	}

	node := soup.HTMLParse(body)
	if node.Error != nil {
		return nil, node.Error
	}

	checkCurrentEvent, err := checkCurrentEventText(node)
	if err != nil {
		return nil, err
	} else if !checkCurrentEvent {
		return nil, domain.ErrNoCurrentEvent
	}

	nodes, err := requiredTraverseHTMLNode(node, currentEventsTraversalOpts)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(nodes); i += 3 {
		format, err := getCurrentEventFormat(nodes[i])
		if err != nil {
			return nil, err
		}

		title, err := getCurrentEventTitle(nodes[i])
		if err != nil {
			return nil, err
		}

		uri, err := getCurrentEventURI(nodes[i])
		if err != nil {
			return nil, err
		}

		date, err := getCurrentEventDate(nodes[i+2])
		if err != nil {
			return nil, err
		}

		duration, err := getCurrentEventDuration(nodes[i+2])
		if err != nil {
			return nil, err
		}

		team, err := getCurrentEventTeam(nodes[i+1])
		if err != nil {
			return nil, err
		}

		upcomingEvents = append(upcomingEvents, domain.CTFTimeEvent{
			Title:    title,
			Format:   format,
			URL:      c.baseURL + uri,
			Date:     date,
			Duration: duration,
			Team:     team,
		})
	}

	return upcomingEvents, nil
}
