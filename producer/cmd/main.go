package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	producer "github.com/frain-dev/webhooks-with-kafka-demo"
	"github.com/segmentio/kafka-go/sasl/scram"

	"github.com/segmentio/kafka-go"
)

func main() {
	c, err := producer.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	w := &kafka.Writer{
		Addr:  kafka.TCP(c.Broker),
		Topic: c.Topic,
	}

	if c.Authentication != nil {
		algo := selectAuthAlgorithm(c.Authentication.Type)
		mechanism, err := scram.Mechanism(algo,
			c.Authentication.Username,
			c.Authentication.Password)

		if err != nil {
			log.Fatal(err)
		}

		w.Transport = &kafka.Transport{
			SASL: mechanism,
			TLS: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	defer func(w *kafka.Writer) {
		err := w.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(w)

	for {
		producer.Produce(context.Background(), w, c.ProduceRate)
		time.Sleep(1 * time.Second)
	}
}

func selectAuthAlgorithm(authType string) scram.Algorithm {
	switch authType {
	case "SHA256":
		return scram.SHA256
	case "SHA512":
		return scram.SHA512
	default:
		return scram.SHA256
	}
}
