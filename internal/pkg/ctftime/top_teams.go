package ctftime

import (
	"fmt"

	"github.com/anaskhan96/soup"
	"github.com/josephsalimin/simple-ctftime-bot/internal/domain"
)

var topTeamsTraversalOpts = []htmlTraversalOption{
	{findType: findOne, findParams: []string{"table", "class", "table table-striped"}},
	{findType: findOne, findParams: []string{"tbody"}},
	{findType: findAll, findParams: []string{"tr"}},
}

var topTeamWorldwidePositionOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 0},
}

var topTeamNameCountryOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 4},
	{findType: findOne, findParams: []string{"a"}},
}

var topTeamPointsCountryOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 5},
}

var topTeamEventsCountryOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 6},
}

var topTeamNameWorldOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 2},
	{findType: findOne, findParams: []string{"a"}},
}

var topTeamPointsWorldOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 4},
}

var topTeamEventsWorldOpts = []htmlTraversalOption{
	{findType: findOneInAll, findParams: []string{"td"}, findIndex: 5},
}

func getTopTeamWorldwidePosition(node soup.Root) (string, error) {
	child, err := requiredTraverseHTMLNode(node, topTeamWorldwidePositionOpts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getTopTeamName(node soup.Root, isWorld bool) (string, error) {
	opts := topTeamNameCountryOpts
	if isWorld {
		opts = topTeamNameWorldOpts
	}

	child, err := requiredTraverseHTMLNode(node, opts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getTopTeamPoints(node soup.Root, isWorld bool) (string, error) {
	opts := topTeamPointsCountryOpts
	if isWorld {
		opts = topTeamPointsWorldOpts
	}

	child, err := requiredTraverseHTMLNode(node, opts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

func getTopTeamEvents(node soup.Root, isWorld bool) (string, error) {
	opts := topTeamEventsCountryOpts
	if isWorld {
		opts = topTeamEventsWorldOpts
	}

	child, err := requiredTraverseHTMLNode(node, opts)
	if err != nil {
		return "", err
	}

	return child[0].Text(), nil
}

// GetTopTeams fetches top team page from CTFTime and parse the HTML to get the value for top team
func (c *Client) GetTopTeams(year int, country string, total int) ([]domain.CTFTimeTeam, error) {
	topTeams := make([]domain.CTFTimeTeam, 0)

	body, err := c.Get(fmt.Sprintf("%v/stats/%v/%v", c.baseURL, year, country))
	if err != nil {
		return nil, err
	}

	root := soup.HTMLParse(body)
	if root.Error != nil {
		return nil, root.Error
	}

	nodes, err := requiredTraverseHTMLNode(root, topTeamsTraversalOpts)
	if err != nil {
		return nil, err
	}

	var size int
	if len(nodes)-1 < total {
		size = len(nodes) - 1
	} else {
		size = total
	}

	isWorld := country == ""
	for i := 1; i <= size; i++ {
		worldwidePosition, err := getTopTeamWorldwidePosition(nodes[i])
		if err != nil {
			return nil, err
		}

		name, err := getTopTeamName(nodes[i], isWorld)
		if err != nil {
			return nil, err
		}

		points, err := getTopTeamPoints(nodes[i], isWorld)
		if err != nil {
			return nil, err
		}

		events, err := getTopTeamEvents(nodes[i], isWorld)
		if err != nil {
			return nil, err
		}

		topTeams = append(topTeams, domain.CTFTimeTeam{
			WorldwidePosition: worldwidePosition,
			Name:              name,
			Points:            points,
			Events:            events,
		})
	}

	return topTeams, nil
}
