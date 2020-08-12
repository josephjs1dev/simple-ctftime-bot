package ctftime

import (
	"github.com/anaskhan96/soup"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
)

var upcomingOpenEventsTraversalOpts = []htmlTraversalOption{
	{findType: findOne, findParams: []string{"div", "id", "upcoming"}},
	{findType: findOne, findParams: []string{"table"}},
	{findType: findOne, findParams: []string{"tbody"}},
	{findType: findAll, findParams: []string{"tr"}},
}

var upcomingEventFormatOpts = []htmlTraversalOption{
	{findType: findOne, findParams: []string{"td", "class", "ctf_format"}},
	{findType: findOne, findParams: []string{"img"}},
}

var upcomingEventTitleOpts = []htmlTraversalOption{
	{findType: findOneInAll, findIndex: 1, findParams: []string{"td"}},
	{findType: findOne, findParams: []string{"a"}},
}

var upcomingEventDateOpts = []htmlTraversalOption{
	{findType: findOneInAll, findIndex: 2, findParams: []string{"td"}},
}

var upcomingEventDurationOpts = []htmlTraversalOption{
	{findType: findOneInAll, findIndex: 3, findParams: []string{"td"}},
}

var upcomingEventTeamOpts = []htmlTraversalOption{
	{findType: findOne, findParams: []string{"small", "class", "muted pull-right"}},
}

func getUpcomingEventFormat(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, upcomingEventFormatOpts)
	if err != nil {
		return "", err
	}

	return getAttrKey(child[0], "title"), nil
}

func getUpcomingEventTitle(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, upcomingEventTitleOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getUpcomingEventURI(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, upcomingEventTitleOpts)
	if err != nil {
		return "", err
	}

	return getAttrKey(child[0], "href"), nil
}

func getUpcomingEventDate(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, upcomingEventDateOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getUpcomingEventDuration(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, upcomingEventDurationOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getUpcomingEventTeam(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, upcomingEventTeamOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

// GetUpcomingEvents fetch CTFTime home page and parse the HTML to get upcoming events data
func (c *Client) GetUpcomingEvents() ([]domain.CTFTimeEvent, error) {
	upcomingEvents := make([]domain.CTFTimeEvent, 0)

	body, err := c.Get(c.baseURL)
	if err != nil {
		return nil, err
	}

	root := soup.HTMLParse(body)
	if root.Error != nil {
		return nil, root.Error
	}

	nodes, err := requiredTraverseHTMLNode(root, upcomingOpenEventsTraversalOpts)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(nodes); i++ {
		format, err := getUpcomingEventFormat(nodes[i])
		if err != nil {
			return nil, err
		}

		title, err := getUpcomingEventTitle(nodes[i])
		if err != nil {
			return nil, err
		}

		uri, err := getUpcomingEventURI(nodes[i])
		if err != nil {
			return nil, err
		}

		date, err := getUpcomingEventDate(nodes[i])
		if err != nil {
			return nil, err
		}

		duration, err := getUpcomingEventDuration(nodes[i])
		if err != nil {
			return nil, err
		}

		team, err := getUpcomingEventTeam(nodes[i])
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
