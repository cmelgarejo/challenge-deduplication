# Backend coding challenge 2023

## Background

When users interact with smart contracts on blockchain, the smart contracts can emit events. It is possible to subscribe to smart contracts events to get real time information about user interactions with the smart contract. In practice we subscribe to a stream of smart contract events by specifying a contract address and a block number. Block number roughly corresponds to a timestamp after which events should be returned.

For this problem we have blockchain events corresponding to payments to our smart contract. We have to process each event not more than once, otherwise we end up with duplicate payments. It is possible that connection to blockchain gets interrupted and when reestablishing connection we receive event that we have processed already. For this reason we need to find a persistent solution to deduplicate events.

## Challenge

Starting github repository <https://github.com/art-technologies/challenge-deduplication>

This repository contains events stored `events.json` that has some duplicate events. Implement persistent solution to deduplicate events in `solution.go` . Event is uniquely identified with `Hash`.

Challenge is to finish `deduplicateEvent` method that can identify duplicate event.

```go
func deduplicateEvent(event *Event) (*Event, error) {
    panic("not implemented")
}
```

## Requirements

- Solution should be persistent under application restarts.
- If event was not processed before then `deduplicateEvent` method should return `event`.
- If event was processed before then `deduplicateEvent` method should return `DuplicateEventError`.
- You can assume that `ProcessEvent` will always succeed and will be executed for deduplicated event.

## Suggestions

For persistence you can use any reliable solution (can try multiple ones). Some ideas for deduplication

- NATS
- Bloom Filter
