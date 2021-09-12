package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Traffic Producer")
}

func gatherStatsFromGet(r *http.Request) []TrafficStats {
	const baseUrl string = "http://server:5000/"
	var url string
	query := r.URL.Query()

	endpointName, epExists := query["endpoint"]

	if !epExists || epExists && len(endpointName) == 0 {
		url = baseUrl + "WeatherForecast"
	} else {
		url = baseUrl + endpointName[0]
	}

	req := Getter{Url: url}
	reqAsync := GetterAsync{Req: req}
	params := TrafficerParams{}
	params.DefaultParams()
	trafficer := Trafficer{}
	stats := trafficer.MakeTraffic(reqAsync, params)

	return stats
}

func rawStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := gatherStatsFromGet(r)

	for _, stat := range stats {
		fmt.Fprintf(w, stat.String())
	}
}

func aggStatsHandler(w http.ResponseWriter, r *http.Request) {
	stats := gatherStatsFromGet(r)
	aggStats := AggregateOnBatch(stats)

	for _, stat := range aggStats {
		fmt.Fprintf(w, stat.String())
	}
}

func getWeatherForecast(w http.ResponseWriter, r *http.Request) {
	const url string = "http://server:5000/WeatherForecast"
	//const url string = "http://localhost:5000/WeatherForecast"

	fmt.Printf("Trying get [%s]...\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "Crap myself while GETing [%s]: %s", url, err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "Cannot read body :(")
	}

	fmt.Fprintf(w, string(body))
	fmt.Fprintf(w, "%d bytes", len(body))
}

func main() {
	port := flag.Int("port", 8080, "Server port")
	flag.Parse()

	http.HandleFunc("/", hello)
	http.HandleFunc("/wf", getWeatherForecast)
	http.HandleFunc("/raw", rawStatsHandler)
	http.HandleFunc("/agg", aggStatsHandler)

	fmt.Printf("Listening on %d...", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
