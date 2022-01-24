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

func (c *EntityNotifier) notify(eventName string, id string, evt []byte) error {
	c.notify(eventName, id, evt)
}
