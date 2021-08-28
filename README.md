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

```
http://localhost:8888/test
```

There is not yet an interface to choose endpoint and control traffic parameters.

