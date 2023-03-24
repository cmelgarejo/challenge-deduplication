package main

import "os"

var processedQueue = make(map[string]bool)

func initNaive() {
	if os.Getenv("SOLUTION") != "naive" {
		return
	}
	// init the map
	processedQueue = make(map[string]bool, len(data))
}

// naiveMapSolution - naive solution with a map, O(1) lookup, O(n) memory
func naiveMapSolution(event *Event) (*Event, error) {
	// if the has exists, raise the error
	if processedQueue[event.Hash] {
		return nil, DuplicateEventError
	} else {
		processedQueue[event.Hash] = true
	}
	return event, nil
}
