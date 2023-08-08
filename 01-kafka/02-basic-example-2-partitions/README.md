1 producer, 1 consumers, 2 partitions

Producer sends messages to 2 partitions. Consumer reads from specific partition which is specified by flag `-partition`.

## Run

#### Docker Compose

```bash
docker-compose up
```

#### Producer

```bash
go run producer/cmd/main.go -count=100 -partition=0
```

#### Consumer

```bash
go run consumer/consumer.go -partition=0
```
