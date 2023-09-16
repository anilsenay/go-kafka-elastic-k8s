FROM golang:latest

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go
EXPOSE 3000
CMD ["./main"]