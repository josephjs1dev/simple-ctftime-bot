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

/** Error Section **/

// ErrNoCurrentEvent ...
var ErrNoCurrentEvent = errors.New("No current event")
