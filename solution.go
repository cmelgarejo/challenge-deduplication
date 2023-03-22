package main

import "errors"

var DuplicateEventError = errors.New("duplicate event")

// deduplicateEvent - checks if the event was processed before.
// Solution should be persisted under application restarts.
func deduplicateEvent(event *Event) (*Event, error) {
	// panic("not implemented")
	return event, nil
}
