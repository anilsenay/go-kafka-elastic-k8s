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

## Minikube

Start minikube with 6GB of memory

```sh
minikube start --memory 7900
```

### Build ElasticSearch

```sh
helm install elasticsearch elastic/elasticsearch -f ./elk-values.yaml
```

Get the password

```sh
kubectl get secrets --namespace=default elasticsearch-master-credentials -ojsonpath='{.data.password}' | base64 -d
```

Port forwarding

```sh
kubectl port-forward svc/elasticsearch-master 9200
```

### Build Kibana

```sh
helm install kibana elastic/kibana
```

Port forwarding

```sh
kubectl port-forward svc/kibana-kibana 5601
```

### Build images

Before building run the following command to use the docker daemon inside minikube:

```sh
eval $(minikube -p minikube docker-env)
```

### Use Elastic Cloud instead of installing ElasticSearch and Kibana

visit here: https://www.elastic.co/cloud

##### Producer

```sh
docker build -t product-producer .
```

##### Consumer

```sh
docker build -t product-consumer .
```

##### Search

```sh
docker build -t search-service .
```

### Apply YAML files

```sh
kubectl apply -f kubernetes/
```

### Port forwarding

```sh
kubectl port-forward svc/product-producer 8080
```

```sh
kubectl port-forward svc/search-service 3000
```

### Create Products

```sh
curl -X POST http://localhost:8080\?count\=5
```

### Check Search

http://localhost:3000/search?keyword=samsung
