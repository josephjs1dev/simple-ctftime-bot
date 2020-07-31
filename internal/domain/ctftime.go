package domain

import "errors"

// CTFTimeEvent ...
type CTFTimeEvent struct {
	Title    string
	URL      string
	Format   string
	Date     string
	Duration string
	Team     string
}

// CTFTimeTeam ...
type CTFTimeTeam struct {
	WorldwidePosition string
	Name              string
	Points            string
	Events            string
}

/** Error Section **/

// ErrNoCurrentEvent ...
var ErrNoCurrentEvent = errors.New("No current event")
