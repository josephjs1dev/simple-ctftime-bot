package ctftime

import (
	"errors"

	"github.com/anaskhan96/soup"
)

// NodeFindType is type for Find function
type NodeFindType int

const (
	findOne NodeFindType = iota + 1
	findAll
	findOneInAll
)

// HTMLDirection ...
type HTMLDirection struct {
	FindType   NodeFindType
	FindIndex  int
	FindParams []string
}

// ErrIndexOutOfRange indicates that function can't find the node from results
var ErrIndexOutOfRange = errors.New("error root FindAll results out of range from FindIndex")

// ErrEmptyResult indicates that function returns empty node
var ErrEmptyResult = errors.New("error empty results from FindAll")

func traverseHTMLNode(root soup.Root, directions []HTMLDirection) ([]soup.Root, error) {
	nodes := []soup.Root{root}
	directionSize := len(directions)

	for _, direction := range directions {
		children := []soup.Root{}
		switch direction.FindType {
		case findOne:
			for _, n := range nodes {
				child := n.Find(direction.FindParams...)
				if child.Error != nil {
					return nil, child.Error
				}

				children = append(children, child)
			}
			nodes = children
		case findAll:
			for idx, n := range nodes {
				res := n.FindAll(direction.FindParams...)
				if len(res) == 0 && idx < directionSize-1 {
					return nil, ErrEmptyResult
				}

				children = append(children, res...)
			}
			nodes = children
		case findOneInAll:
			for _, n := range nodes {
				res := n.FindAll(direction.FindParams...)
				if len(res) < direction.FindIndex {
					return nil, ErrIndexOutOfRange
				}

				children = append(children, res[direction.FindIndex])
			}
			nodes = children
		}
	}

	return nodes, nil
}

func requiredTraverseHTMLNode(root soup.Root, directions []HTMLDirection) ([]soup.Root, error) {
	nodes, err := traverseHTMLNode(root, directions)
	if err != nil {
		return nil, err
	}

	if len(nodes) == 0 {
		return nil, ErrEmptyResult
	}

	return nodes, nil
}

func getAttrKey(node soup.Root, key string) string {
	attr := node.Attrs()
	if attr != nil {
		v, ok := attr["href"]
		if ok {
			return v
		}
	}

	return ""
}
