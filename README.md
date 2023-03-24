# Backend coding challenge 2023

Hi Augus! I hope you're doing well. I'm sending you this challenge as a part of the interview process for the backend developer position. I'm really excited about this opportunity! I hope you'll like my solution anyway. :)

Some notes:

I'm using vscode so I've added the launch.json and some env vars to use with it, using docker to setup the persistance layer (postgres in this case)

I wanted to add NATS too but I haven't had the time for reading more about JetStream, also I thought about
put in temportal.io and NATS but I didn't have the time to implement them too.

I was thinking of:

    - Setting NATS in JetStream mode and use message deduplication as shown [here](https://docs.nats.io/using-nats/developer/develop_jetstream/model_deep_dive#message-deduplication)
    - Using temporal.io to schedule concurrent event load and send them to NATS

I haven't touched anything else than solution.go, wanted to add some benchmarks on main.go, but basically its just injecting everything through the solution.go file.

## Running the solution

    ```bash
    make setup #give it a sec, so the pgsql db spins up, then...
    make run
    ```

## Solutions

Implemented a couple of solution for the task, using different approaches.

### Solution 1

Naive map based solution. It's not optimal at all, but it's simple and fast.

### Solution 2

Using bloom filter implementation, that has been shamelessly ripped from [here](https://medium.com/the-little-bit-ninja/bloom-filters-and-when-to-use-them-ab64028996d4)

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
