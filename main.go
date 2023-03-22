package main

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
)

//go:embed events.json
var data []byte

func steamEvents(afterBlockNumber int) []Event {
	events := make([]Event, 0)
	_ = json.Unmarshal(data, &events)

	filteredEvents := make([]Event, 0)

	for i := range events {
		if events[i].Block >= afterBlockNumber {
			filteredEvents = append(filteredEvents, events[i])
		}
	}

	return filteredEvents
}

func main() {
	events := steamEvents(0)

	for i := range events {
		event, err := deduplicateEvent(&events[i])
		if errors.Is(err, DuplicateEventError) {
			fmt.Println("duplicate found", events[i])
			continue
		} else if err != nil {
			panic(err)
		}

		ProcessEvent(event)
	}
}
