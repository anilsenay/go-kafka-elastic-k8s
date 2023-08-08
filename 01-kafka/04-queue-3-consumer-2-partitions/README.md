1 producer, 3 consumers, 2 partitions

Producer sends messages to 2 partitions. Each consumer TRY to read from one partition without specify partition. But only one consumer can read from each partition. So last consumer will not read any message.

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
