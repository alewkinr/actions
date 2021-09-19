FROM golang:1.16 as builder

WORKDIR /app
COPY . /app

RUN CGO_ENABLED=0 go build -o bin/service ./cmd/main.go

FROM alpine:3.14

COPY --from=builder /app/bin/service /service
CMD ["/service"]