package main

import (
	"events-dedup/internal/lib/bloom"
	"fmt"
	"os"
)

var bloomFilter *bloom.BloomFilter

func initBloom() {
	// skip if the solution is not "bloom"
	if os.Getenv("SOLUTION") != "bloom" {
		return
	}
	count := len(steamEvents(0))
	// 1.08 is the sweetspot to generate an optimal length for this particular case,
	// as 1.09 gives an 0.072568940493469% false positive rate
	falsePositivePercentage := 1.08
	optimalLength, optimalRounds := bloom.CalculateOptimalParameters(count, falsePositivePercentage)
	fmt.Printf("Item count: %d. Target false positive percentage: %.2f. Optimal length: %d. Optimal rounds: %d\n",
		count, falsePositivePercentage, optimalLength, optimalRounds)
	bloomFilter = bloom.NewBloomFilter(optimalLength, optimalRounds, bloom.HashFuncSha)
}

func bloomSolution(event *Event) (*Event, error) {
	if bloomFilter.IsMember([]byte(event.Hash)) {
		return nil, DuplicateEventError
	}
	bloomFilter.Add([]byte(event.Hash))
	return event, nil
}
