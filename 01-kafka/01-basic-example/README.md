1 producer, 1 consumers, 1 partitions

Producer sends messages to 1 partitions. Consumer reads from that partition.

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
go run consumer/consumer.go
```
