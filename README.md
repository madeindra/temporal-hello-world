# Hello World with Temporal

## What is Temporal?

Temporal is workflow orchestrator, it ensures success in code execution, it helps for non-deterministic algorithm.

## What are the parts of Temporal?

* Workflow: Activity wrapper to be able to run in Temporal
* Activity: Use cases or logic that want to be executed
* Worker: Workflow runner / retyer / scheduler
* Initiator (start): Trigger workflow execution

## Requirements
* Docker
* Docker Compose
* Temporal SDK (Go)

## Start Temporal

1. Clone temporal docker repo
```
git clone https://github.com/temporalio/docker-compose.git
```

2. Start Temporal
```
cd docker-compose
docker-compose up
```

## Get the Go SDK

1. Get the SDK
```
go get go.temporal.io/sdk
```

2. or clone the Go SDK repo
```
git clone https://github.com/temporalio/sdk-go.git
```

## Run Application

1. Start the worker
```
go run worker/main.go
```

2. Start the initiator
```
go run start/main.go
```

## References

This project is created based on tutorial on [Temporal](https://learn.temporal.io/getting_started/go/hello_world_in_go/).