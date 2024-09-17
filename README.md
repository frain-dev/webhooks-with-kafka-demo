# webhooks-with-kafka-demo
This repository includes all the artefacts used to demonstrate ingesting and transforming events from Kafka into Convoy to send out webhooks.

## Prerequisites
1. Convoy Instance
2. Kafka Cluster

## Instructions
```bash
# Produce events to Kafka
$ go run ./producer/cmd
```
