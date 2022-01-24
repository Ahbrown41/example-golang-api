package models

import (
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
)

const dbName = "connection_test.db"

func TestMain(m *testing.M) {
	//setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestConnectDB(t *testing.T) {
	dialect := sqlite.Open(dbName)
	conn := New(dialect)
	assert.NotNil(t, conn)
	err := conn.Connect()
	assert.Nil(t, err)
	defer conn.Disconnect()
}

func shutdown() {
	e := os.Remove(dbName)
	if e != nil {
		log.Err(e)
	}
}
