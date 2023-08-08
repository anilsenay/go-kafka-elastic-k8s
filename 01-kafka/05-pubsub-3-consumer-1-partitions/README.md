1 producer, 3 consumers, 3 consumer group, 1 partitions

Producer sends messages to 1 partitions. Each consumer reads same data from that partition.

## Run

#### Docker Compose

```bash
docker-compose up
```

#### Producer

```bash
go run producer/cmd/main.go -count=100
```

#### Consumer 1

```bash
go run consumer-0/consumer.go
```

#### Consumer 2

```bash
go run consumer-1/consumer.go
```

#### Consumer 3

```bash
go run consumer-2/consumer.go
```
