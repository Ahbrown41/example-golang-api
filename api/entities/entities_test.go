package entities

import (
	"api-reference-golang/api/utils"
	"api-reference-golang/models"
	"api-reference-golang/models/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"os"
	"testing"
	"time"
)

const baseURL = "/api/v1/entities"
const dbName = "entities_test.db"

var (
	router *gin.Engine
	db     *models.Connection
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func setup() {
	// Open DB
	dialect := sqlite.Open(dbName)
	db = models.New(dialect)
	err := db.Connect()
	if err != nil {
		log.Err(err)
	}

	// Create Routes
	router = gin.Default()
	CreateRoutes(router, db)

	// Migrate what is necessary for test
	entity.Migrate(db.DB())
}

func shutdown() {
	db.Disconnect()
	e := os.Remove(dbName)
	if e != nil {
		log.Err(e)
	}
}

func TestCreateEntity(t *testing.T) {
	ety := entity.Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
		Items: []entity.EntityItem{entity.EntityItem{ItemName: "#1"}, entity.EntityItem{ItemName: "#2"}},
	}
	url := fmt.Sprintf("%s/", baseURL)
	rr := utils.ExecuteReq(t, router, "POST", url, ety)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var entity entity.Entity
	err := json.Unmarshal(rr.Body.Bytes(), &entity)
	assert.Nil(t, err)
	assert.NotNil(t, entity, "Entity is nil")
	assert.Greater(t, entity.ID, uint(0), "Entity is nil")
}

func TestAllEntities(t *testing.T) {
	url := fmt.Sprintf("%s/", baseURL)
	rr := utils.ExecuteReq(t, router, "GET", url, nil)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var entities []entity.Entity
	err := json.Unmarshal(rr.Body.Bytes(), &entities)
	assert.Nil(t, err)
	assert.NotNil(t, entities, "Entity is nil")
	assert.Greater(t, len(entities), 0, "Entity is nil")
}

func TestGetEntityByID(t *testing.T) {
	url := fmt.Sprintf("%s/", baseURL)
	rr := utils.ExecuteReq(t, router, "GET", url, nil)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var entities []entity.Entity
	err := json.Unmarshal(rr.Body.Bytes(), &entities)
	assert.Nil(t, err)
	assert.NotNil(t, entities, "Entity is nil")
	assert.Greater(t, len(entities), 0, "Entity is nil")

	for _, k := range entities {
		url := fmt.Sprintf("%s/%d", baseURL, k.ID)
		rr = utils.ExecuteReq(t, router, "GET", url, nil)
		assert.Equal(t, 200, rr.Code, "Return code not 200")
		var entity entity.Entity
		err := json.Unmarshal(rr.Body.Bytes(), &entity)
		assert.Nil(t, err)
		assert.Greater(t, entity.ID, uint(0), "Entity is nil")
		break
	}
}

func TestUpdateEntity(t *testing.T) {
	// Create Item
	ety := entity.Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
		Items: []entity.EntityItem{entity.EntityItem{ItemName: "#1"}, entity.EntityItem{ItemName: "#2"}},
	}
	url := fmt.Sprintf("%s/", baseURL)
	rr := utils.ExecuteReq(t, router, "POST", url, ety)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var entity entity.Entity
	err := json.Unmarshal(rr.Body.Bytes(), &entity)
	assert.Nil(t, err)
	assert.NotNil(t, entity, "Entity is nil")
	assert.Greater(t, entity.ID, uint(0), "Entity is nil")

	// Update Item
	url = fmt.Sprintf("%s/%d", baseURL, entity.ID)
	rr = utils.ExecuteReq(t, router, "PUT", url, ety)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	err = json.Unmarshal(rr.Body.Bytes(), &entity)
	assert.Nil(t, err)
	assert.NotNil(t, entity, "Entity is nil")
	assert.Greater(t, entity.ID, uint(0), "Entity is nil")
}

func TestDeleteEntity(t *testing.T) {
	// Create Item
	ety := entity.Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
		Items: []entity.EntityItem{entity.EntityItem{ItemName: "#1"}, entity.EntityItem{ItemName: "#2"}},
	}
	url := fmt.Sprintf("%s/", baseURL)
	rr := utils.ExecuteReq(t, router, "POST", url, ety)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var entity entity.Entity
	err := json.Unmarshal(rr.Body.Bytes(), &entity)
	assert.Nil(t, err)
	assert.NotNil(t, entity, "Entity is nil")
	assert.Greater(t, entity.ID, uint(0), "Entity has 0 ID: %v", entity)

	// Delete Item
	url = fmt.Sprintf("%s/%d", baseURL, entity.ID)
	rr = utils.ExecuteReq(t, router, "DELETE", url, ety)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	err = json.Unmarshal(rr.Body.Bytes(), &entity)
	assert.Nil(t, err)
	assert.NotNil(t, entity, "Entity is nil")
	assert.Greater(t, entity.ID, uint(0), "Entity has 0 ID")
}
