package main

import "sort"

type AggBatch struct {
	BatchId            int
	Count              int
	ResponseBytes50q   float64
	ResponseTimeMsg50q float64
	ResponseBytes95q   float64
	ResponseTimeMs95q  float64
	Errors             int
}

func AggregateOnBatch(stats []TrafficStats) []AggBatch {
	agg := make([]AggBatch, 0)

	sortTrafficStats(stats)
	statsBatches := filterBatches(stats)

	for batchId := range statsBatches {
		byteResponses := selectResponseBytesAsFloats(statsBatches[batchId])
		timeResponses := selectResponseTimeMsAsFloats(statsBatches[batchId])

		cnt := len(statsBatches[batchId])
		bytesRespQuantiles := Quantiles(byteResponses, []float64{0.50, 0.95})
		timeRespQuantiles := Quantiles(timeResponses, []float64{0.50, 0.95})

		aggCurr := AggBatch{
			BatchId:            batchId,
			Count:              cnt,
			ResponseBytes50q:   bytesRespQuantiles[0],
			ResponseBytes95q:   bytesRespQuantiles[1],
			ResponseTimeMsg50q: timeRespQuantiles[0],
			ResponseTimeMs95q:  timeRespQuantiles[1],
		}

		agg = append(agg, aggCurr)
	}

	return agg
}

// This method assumes that stats are already sorted by (BatchId, RequestId)
// TODO: Test filtering
func filterBatches(stats []TrafficStats) map[int][]TrafficStats {
	batchIdToStats := make(map[int][]TrafficStats)

	if len(stats) == 0 {
		return batchIdToStats
	}

	prevBatchId := stats[0].BatchId
	prevId := 0

	for id, stat := range stats {
		if stat.BatchId != prevBatchId {
			batchIdToStats[prevBatchId] = stats[prevId:id]
			prevId = id
			prevBatchId = stat.BatchId
		}

		if id == len(stats)-1 {
			batchIdToStats[prevBatchId] = stats[prevId:id]
		}
	}

	return batchIdToStats
}

func sortTrafficStats(stats []TrafficStats) {
	sort.SliceStable(stats, func(i, j int) bool {
		if stats[i].BatchId == stats[j].BatchId {
			return stats[i].RequestId < stats[j].RequestId
		}
		return stats[i].BatchId < stats[j].BatchId
	})
}

func selectResponseBytesAsFloats(stats []TrafficStats) []float64 {
	res := make([]float64, len(stats))

	for id, e := range stats {
		res[id] = float64(e.ResponseBytes)
	}

	return res
}

func selectResponseTimeMsAsFloats(stats []TrafficStats) []float64 {
	res := make([]float64, len(stats))

	for id, e := range stats {
		res[id] = float64(e.ResponseTimeMs)
	}

	return res
}
