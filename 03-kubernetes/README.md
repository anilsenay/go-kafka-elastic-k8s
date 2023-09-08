## Run

### Docker Compose

```bash
docker-compose up
```

### Producer

```bash
go run producer/cmd/main.go -count=100
```

### Consumer

```bash
go run consumer-0/consumer.go
```

## Search

```bash
go run search/cmd/main.go
```

Go to http://localhost:3000/search
