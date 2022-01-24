package entity

import (
	"api-reference-golang/events/kafka"
)

type EntityNotifier struct {
	producer *kafka.Producer
}

func NewEntityNotifier(producer *kafka.Producer) EntityNotifier {
	return EntityNotifier{producer: producer}
}

func (c *EntityNotifier) Notify(eventName string, id string, event []byte) error {
	return c.producer.Notify(eventName, id, event)
}
