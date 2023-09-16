FROM golang:1.21.1-alpine

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev librdkafka-dev pkgconf

WORKDIR /app
COPY . /app
RUN go build -tags musl -o main cmd/main.go
EXPOSE 8080
CMD ["./main"]