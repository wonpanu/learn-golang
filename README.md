# REST API with Go, mongoDB and Elasticsearch

### TL;DR

Implement REST API using ``` Go + Fiber + MongoDB + Elasticsearch ```

### Diagram Design

![alt text](https://github.com/wonpanu/learn-golang/blob/main/Workflow.png?raw=true)

### To run MongoDB docker

```
docker-compose up -d
```

### To run service

##### Please run docker before run service

```
cd service
go run cmd/main.go
```

### To run worker

##### Please run docker before run worker

```
cd worker/bulk/cmd
go run main.go
```