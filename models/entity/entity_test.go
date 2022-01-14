package entity

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"os"
	"reference-golang-svc/models"
	"testing"
	"time"
)

var db *models.Connection

func TestMain(t *testing.M) {
	setup()
	code := t.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	dialect := sqlite.Open("entity_test.db")
	conn := models.New(dialect)
	db = conn
	db.Connect()

	// Migrate what is necessary for test
	db.DB().AutoMigrate(
		Entity{},
	)
}

func shutdown() {
	defer db.Disconnect()
}

func CreateEntity(t *testing.T) {
	svc := New(db.DB())
	entities1, err := svc.All(FetchOptions{Page: 1, Limit: 10})
	assert.Nil(t, err)
	id, err := svc.Create(Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Time{},
	})
	assert.Nil(t, err)
	assert.Greater(t, id, 0, "New ID ==0")
	entities2, err := svc.All(FetchOptions{Page: 1, Limit: 10})
	assert.Nil(t, err)
	assert.Greater(t, len(*entities2), len(*entities1), "Only 0 entities")
}

func TestAll(t *testing.T) {
	svc := New(db.DB())
	svc.Create(Entity{
		Name:  "Entity 1",
		Value: 43,
		Date:  time.Time{},
	})
	entities, err := svc.All(FetchOptions{Page: 1, Limit: 10})
	assert.Nil(t, err)
	assert.Greater(t, len(*entities), 0, "Only 0 entities")
}
