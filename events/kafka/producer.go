package kafka

import (
	"context"
	json2 "encoding/json"
	cloudevents "github.com/cloudevents/sdk-go/v2/event"
	"github.com/rs/zerolog/log"
	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	write     *kafka.Writer
	source    string
	eventType string
}

// New - Creates kafka
func New(addr string, topic string, source string, eventType string) *Producer {
	write := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &Producer{write: write, source: source, eventType: eventType}
}

// Disconnect - Disconnects kafka connection
func (c *Producer) Disconnect() {
	if err := c.write.Close(); err != nil {
		log.Fatal().Msgf("failed to close writer: %s", err)
	}
}

// Produce - Produces Kafka Message
func (c *Producer) Notify(eventName string, id string, evt []byte) error {
	// Build CloudEvent
	event := cloudevents.New()
	event.SetID(id)
	event.SetSource(c.source)
	event.SetType(c.eventType)
	event.SetData(cloudevents.ApplicationJSON, evt)

	json, err := json2.Marshal(event)

	if err != nil {
		log.Error().Err(err)
		return err
	}
	// Send Message
	msg := kafka.Message{
		Key:   []byte(id),
		Value: json,
	}
	err = c.write.WriteMessages(context.Background(),
		msg,
	)
	if err != nil {
		log.Error().Msgf("failed to write messages: %s", err)
		return err
	}
	return nil
}
