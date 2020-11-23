FROM golang:1.12 as builder

WORKDIR /app
COPY . .

RUN go get github.com/prometheus/client_golang/prometheus github.com/prometheus/client_golang/prometheus/promhttp &&\
    CGO_ENABLED=0 GOOS=linux go build -v -o exporter

FROM alpine
RUN apk add --no-cache ca-certificates

COPY --from=builder /app/exporter /exporter

CMD ["/exporter"]