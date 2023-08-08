docker compose exec kafka1 \
  kafka-topics --create \
    --topic mytopic \
    --bootstrap-server localhost:9092 \
    --replication-factor 1 \
    --partitions 1