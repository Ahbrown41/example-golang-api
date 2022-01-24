package entity

import (
	"api-reference-golang/events/kafka"
	"api-reference-golang/models"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
	"time"
)

var (
	db  *models.Connection
	svc Repository
)

const dbName = "entity_test.db"

func TestMain(t *testing.M) {
	setup()
	code := t.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	dialect := sqlite.Open(dbName)
	conn := models.New(dialect)
	db = conn
	db.Connect()

	kafka := kafka.New("localhost:9092", "entities", "api-reference-golang")

	// Migrate what is necessary for test
	db.DB().AutoMigrate(
		Entity{},
		EntityItem{},
	)

	svc = New(db, kafka)
}

func shutdown() {
	db.Disconnect()
	e := os.Remove(dbName)
	if e != nil {
		log.Err(e)
	}
}

func TestCreateEntity(t *testing.T) {
	entities1, err := svc.All(FetchOptions{Page: 1, Limit: 10})
	assert.Nil(t, err)
	entity, err := svc.Create(&Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
	})
	assert.Nil(t, err)
	assert.Greater(t, entity.ID, uint(0), "New ID ==0")
	entities2, err := svc.All(FetchOptions{Page: 1, Limit: 10})
	assert.Nil(t, err)
	assert.Greater(t, len(*entities2), len(*entities1), "Only 0 entities")
}

func TestCreateEntityAndItems(t *testing.T) {
	entity, err := svc.Create(&Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
		Items: []EntityItem{EntityItem{ItemName: "#1"}, EntityItem{ItemName: "#2"}},
	})
	assert.Nil(t, err)
	assert.Greater(t, entity.ID, uint(0), "New ID ==0")
	entity2, err := svc.Get(entity.ID)
	assert.Equal(t, 2, len(entity2.Items), "Entity Items should be 2")
}

func TestAll(t *testing.T) {
	entity, err := svc.Create(&Entity{
		Name:  "Entity 1",
		Value: 43,
		Date:  time.Now(),
	})
	assert.Nil(t, err)
	assert.Greater(t, entity.ID, uint(0), "Entity ID == 0")
	entities, err := svc.All(FetchOptions{Page: 1, Limit: 10})
	assert.Nil(t, err)
	assert.Greater(t, len(*entities), 0, "Only 0 entities")
}

func TestCreateEntityValidationError(t *testing.T) {
	entity, err := svc.Create(&Entity{
		Name:  "Entity X",
		Value: 150,         //Invalid > 100
		Date:  time.Time{}, //Invalid, empty date
	})
	assert.NotNil(t, err)
	assert.Nil(t, entity, "Entity ID not nil")
	errors := models.Errors{}
	json.Unmarshal([]byte(err.Error()), &errors)
	assert.Equalf(t, "*entity.Entity", errors.Struct, "Validated Struc invalid")
	assert.Greater(t, len(errors.Errors), 0, "Errors not > 0")
}
