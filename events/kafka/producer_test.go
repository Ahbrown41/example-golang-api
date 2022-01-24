package kafka

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

var (
	prod *Producer
)

type Test struct {
	id    uint
	name  string
	value string
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	prod = New("localhost:9092", "topic-A", "producer_test")
}

func shutdown() {
	prod.Disconnect()
}

func TestProduceMsg(t *testing.T) {
	msg := Test{id: uint(123), name: "Test", value: "TestVal"}
	byt, _ := json.Marshal(msg)
	err := prod.Notify("order-1", strconv.FormatUint(uint64(msg.id), 10), byt)
	assert.Nil(t, err)
}
