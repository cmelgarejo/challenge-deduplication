package main

import "fmt"

// Event - describes payment on blockchain
type Event struct {
	Hash      string `json:"hash"`
	Block     int    `json:"block"`
	Timestamp int    `json:"timestamp"`
	Amount    string `json:"amount"`
	From      string `json:"from"`
}

// ProcessEvent - business logic to process received event and assign it to the user.
// This function takes deduplicated event.
// This function does not need be implemented.
func ProcessEvent(event *Event) {
	fmt.Println(event)
}
