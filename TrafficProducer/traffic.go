package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Trafficer struct{}

type TrafficStats struct {
	BatchId        int
	RequestId      int
	ResponseBytes  int
	ResponseTimeMs int64
	IsError        bool
}

type TrafficerParams struct {
	RequestsAtOnce int
	PauseLength    time.Duration
	Batches        int
}

func (t *Trafficer) MakeTraffic(requester RequesterAsync, params TrafficerParams) []TrafficStats {
	n := params.Batches * params.RequestsAtOnce
	stats := make([]TrafficStats, 0, n)
	statsChan := make(chan TrafficStats, n)
	var wg sync.WaitGroup

	for batchId := 0; batchId < params.Batches; batchId++ {
		log.Printf("Starting batch [%d]\n", batchId)

		for requestId := 0; requestId < params.RequestsAtOnce; requestId++ {
			wg.Add(1)
			go requester.RequestAsync(batchId, requestId, statsChan, &wg)
		}
		wg.Wait()
	}
	close(statsChan)

	for stat := range statsChan {
		stats = append(stats, stat)
	}

	return stats
}

// Setup default parameters.
func (tp *TrafficerParams) DefaultParams() {
	tp.RequestsAtOnce = 100
	tp.PauseLength = time.Second * 5
	tp.Batches = 10
}

func (ts *TrafficStats) String() string {
	return fmt.Sprintf("[%d | %d] Time [ms]: %d | Bytes: %d | IsError: %t\n",
		ts.BatchId, ts.RequestId, ts.ResponseTimeMs, ts.ResponseBytes, ts.IsError)
}

// Prepare parameters
func MakeParams(requests int, pause time.Duration, batches int) TrafficerParams {
	return TrafficerParams{
		RequestsAtOnce: requests,
		PauseLength:    pause,
		Batches:        batches,
	}
}
