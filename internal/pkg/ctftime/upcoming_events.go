package ctftime

import (
	"github.com/anaskhan96/soup"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
)

var upcomingOpenEventsDirections = []HTMLDirection{
	{FindType: findOne, FindParams: []string{"div", "id", "upcoming"}},
	{FindType: findOne, FindParams: []string{"table"}},
	{FindType: findOne, FindParams: []string{"tbody"}},
	{FindType: findAll, FindParams: []string{"tr"}},
}

var eventCTFFormat = []HTMLDirection{
	{FindType: findOne, FindParams: []string{"td", "class", "ctf_format"}},
	{FindType: findOne, FindParams: []string{"img"}},
}

var eventCTFTitle = []HTMLDirection{
	{FindType: findOneInAll, FindIndex: 1, FindParams: []string{"td"}},
	{FindType: findOne, FindParams: []string{"a"}},
}

var eventCTFDate = []HTMLDirection{
	{FindType: findOneInAll, FindIndex: 2, FindParams: []string{"td"}},
}

var eventCTFDuration = []HTMLDirection{
	{FindType: findOneInAll, FindIndex: 3, FindParams: []string{"td"}},
}

func getCTFFormat(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, eventCTFFormat)
	if err != nil {
		return "", err
	}

	return getAttrKey(child[0], "data-original-title"), nil
}

func getCTFTitle(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, eventCTFTitle)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getCTFURI(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, eventCTFTitle)
	if err != nil {
		return "", err
	}

	return getAttrKey(child[0], "href"), nil
}

func getCTFDate(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, eventCTFDate)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getCTFDuration(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, eventCTFDuration)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

// GetUpcomingEvents ...
func (c *Client) GetUpcomingEvents() ([]domain.CTFTimeUpcomingEvent, error) {
	upcomingEvents := make([]domain.CTFTimeUpcomingEvent, 0)

	body, err := c.Get(c.baseURL)
	if err != nil {
		return nil, err
	}

	node := soup.HTMLParse(body)

	if node.Error != nil {
		return nil, node.Error
	}

	nodes, err := requiredTraverseHTMLNode(node, upcomingOpenEventsDirections)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(nodes); i++ {
		format, err := getCTFFormat(nodes[i])
		if err != nil {
			return nil, err
		}

		title, err := getCTFTitle(nodes[i])
		if err != nil {
			return nil, err
		}

		uri, err := getCTFURI(nodes[i])
		if err != nil {
			return nil, err
		}

		date, err := getCTFDate(nodes[i])
		if err != nil {
			return nil, err
		}

		duration, err := getCTFDuration(nodes[i])
		if err != nil {
			return nil, err
		}

		upcomingEvents = append(upcomingEvents, domain.CTFTimeUpcomingEvent{
			Title:    title,
			Format:   format,
			URL:      c.baseURL + uri,
			Date:     date,
			Duration: duration,
		})
	}

	return upcomingEvents, nil
}
