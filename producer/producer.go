package producer

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"

	"github.com/jaswdr/faker/v2"
	"github.com/segmentio/kafka-go"
)

var eventTypes, businessIds []string

func init() {
	// Initialize event types
	eventTypes = []string{"invoice.created", "payment.created"}

	// Seed Business IDs
	businessIds = []string{
		"2f8f7425-5fb7-422f-b5d5-da7911a0e1a5",
		"f5f19855-b8da-4b7a-9154-e6093765bb17"}
}

// Event is the payload on Kafka.
type Event struct {
	ID         string `json:"id"`
	EventType  string `json:"event_type"`
	BusinessId string `json:"business_id"`
	Email      string `json:"email"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Address    string `json:"address"`
	Amount     int    `json:"amount"`
}

// Produce generates random events to the configured Kafka topic.
func Produce(ctx context.Context, w *kafka.Writer, produceRate uint) {
	var messages []kafka.Message

	for i := 0; i < int(produceRate); i++ {
		payload, err := generatePayload()
		if err != nil {
			log.Fatal(err)
		}

		messages = append(messages, kafka.Message{Value: payload})
	}

	err := w.WriteMessages(ctx, messages...)
	if err != nil {
		log.Fatal(err)
	}
}

func generatePayload() ([]byte, error) {
	fake := faker.New()

	fakeEvent := &Event{
		ID:         fake.UUID().V4(),
		EventType:  eventTypes[rand.Intn(2)],
		BusinessId: businessIds[rand.Intn(2)],
		Email:      fake.Internet().Email(),
		FirstName:  fake.Person().FirstName(),
		LastName:   fake.Person().LastName(),
		Address:    fake.Address().Address(),
		Amount:     fake.Currency().Number(),
	}

	payload, err := json.Marshal(fakeEvent)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
