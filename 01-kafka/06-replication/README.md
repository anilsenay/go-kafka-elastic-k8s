1 producer, 1 consumers, 1 partitions, 3 replicas

Producer sends messages to 1 partitions. Consumer reads from that partition. 3 replicas.

## Run

#### Docker Compose

```bash
docker-compose up
```

#### Producer

```bash
go run producer/cmd/main.go -count=10000
```

#### Consumer

```bash
go run consumer/consumer.go
```

#### Check who is the leader

```bash
./get-leader.sh
```

#### Stop one of the brokers

```bash
docker stop kafka-2
```

#### Produce more messages

```bash
go run producer/cmd/main.go -count=10000
```

#### Start the broker again

```bash
docker start kafka-2
```

#### Stop other brokers

```bash
docker stop kafka-1
docker stop kafka-3
```

#### Check who is the leader again (it should be kafka-2)

```bash
./get-leader.sh
```

#### Consume messages while only kafka-2 is running

```bash
go run consumer/consumer.go
```
