docker compose exec kafka3 \
  kafka-topics --create \
    --topic mytopic \
    --bootstrap-server localhost:9094 \
    --replication-factor 3 \
    --partitions 1