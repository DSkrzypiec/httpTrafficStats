package main

import (
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Requester interface {
	Request() ([]byte, error)
}

type Getter struct {
	Url string
}

func (g Getter) Request() ([]byte, error) {
	resp, err := http.Get(g.Url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

type RequesterAsync interface {
	RequestAsync(batchId, requestId int,
		stats chan TrafficStats, wg *sync.WaitGroup)
}

type GetterAsync struct {
	Req Requester
}

func (ga GetterAsync) RequestAsync(batchId, requestId int,
	stats chan TrafficStats, wg *sync.WaitGroup) {

	start := time.Now()
	bytes, err := ga.Req.Request()
	elapsedMs := time.Since(start).Milliseconds()

	currStats := TrafficStats{
		BatchId:        batchId,
		RequestId:      requestId,
		ResponseBytes:  len(bytes),
		ResponseTimeMs: elapsedMs,
		IsError:        false,
	}

	if err != nil {
		currStats.IsError = true
	}

	stats <- currStats
	wg.Done()
}
