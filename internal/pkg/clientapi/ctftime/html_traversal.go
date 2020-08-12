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

type htmlTraversalOption struct {
	findType   NodeFindType
	findIndex  int
	findParams []string
}

// ErrIndexOutOfRange indicates that function can't find the node from results
var ErrIndexOutOfRange = errors.New("error root FindAll results out of range from FindIndex")

// ErrEmptyResult indicates that function returns empty node
var ErrEmptyResult = errors.New("error empty results from FindAll")

func traverseHTMLNode(root soup.Root, opts []htmlTraversalOption) ([]soup.Root, error) {
	nodes := []soup.Root{root}
	optSize := len(opts)

	for _, opt := range opts {
		children := []soup.Root{}
		switch opt.findType {
		case findOne:
			for _, n := range nodes {
				child := n.FindStrict(opt.findParams...)
				if child.Error != nil {
					return nil, child.Error
				}

				children = append(children, child)
			}
			nodes = children
		case findAll:
			for idx, n := range nodes {
				res := n.FindAllStrict(opt.findParams...)
				if len(res) == 0 && idx < optSize-1 {
					return nil, ErrEmptyResult
				}

				children = append(children, res...)
			}
			nodes = children
		case findOneInAll:
			for _, n := range nodes {
				res := n.FindAllStrict(opt.findParams...)
				if len(res) <= opt.findIndex {
					return nil, ErrIndexOutOfRange
				}

				children = append(children, res[opt.findIndex])
			}
			nodes = children
		}
	}

	return nodes, nil
}

func requiredTraverseHTMLNode(root soup.Root, opts []htmlTraversalOption) ([]soup.Root, error) {
	nodes, err := traverseHTMLNode(root, opts)
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
		v, ok := attr[key]
		if ok {
			return v
		}
	}

	return ""
}
