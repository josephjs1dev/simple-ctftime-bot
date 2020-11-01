package ctftime

import "errors"

const (
	baseURL = "https://ctftime.org"
)

// ErrNoCurrentEvent indicates that it can't find current event.
var ErrNoCurrentEvent = errors.New("no current event")

// ErrIndexOutOfRange indicates that function can't find the node from results.
var ErrIndexOutOfRange = errors.New("root FindAll results out of range from FindIndex")

// ErrEmptyResult indicates that function returns empty node.
var ErrEmptyResult = errors.New("empty results from FindAll")

// Event ...
type Event struct {
	Title    string
	URL      string
	Format   string
	Date     string
	Duration string
	Team     string
}

// Team ...
type Team struct {
	Name              string
	Points            string
	Events            string
	WorldwidePosition string
}
