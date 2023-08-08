1 producer, 2 consumers, 2 partitions

Producer sends messages to 2 partitions. Each consumer reads from one partition.

## Run

#### Docker Compose

```bash
docker-compose up
```

#### Producer

```bash
go run producer/cmd/main.go -count=100
```

#### Consumer

```bash
go run consumer-0/consumer.go
```
