# HTTP Traffic Statistics

This project is a little tool to quickly test and measure your HTTP server endpoints.
It's based on `docker-compose` network therefore latency due to actual netowork connection is ommited.

There are two components: HTTP server (target) and Traffic Producer (also HTTP server).
Traffic Producer would send many requests to the HTTP server and gather stats about responses.


## Run

```
docker-compose build
docker-compose up
```

To get raw (not aggregated) results you can use `/raw` endpoint of Traffic Producer. To specify
which `[GET]` endpoint of the HTTP server just provide `endpoint` URL query value:

```
http://localhost:8888/raw?endpoint=test
```

To get statistics (aggregated results) you can use `/agg` endpoint of Traffic Producer. To specify
which `[GET]` endpoint of the HTTP server just provide `endpoint` URL query value:

```
http://localhost:8888/agg?endpoint=test2
```

There is not yet an interface to control traffic parameters.

