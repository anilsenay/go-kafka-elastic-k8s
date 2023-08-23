docker compose exec kafka3 \
  kafka-topics --bootstrap-server localhost:9094 --describe --topic mytopic
