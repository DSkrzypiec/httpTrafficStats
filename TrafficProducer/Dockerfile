FROM golang:1.16

WORKDIR /app
COPY . .
RUN go build -o TrafficProducer
ENTRYPOINT ["/app/TrafficProducer", "-port", "8080"]

