package entities

import (
	"api-reference-golang/models"
	"api-reference-golang/models/entity"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	router = gin.Default()
	CreateRoutes(router)

	// Open DB
	dialect := sqlite.Open(dbName)
	db = models.New(dialect)
	err := db.Connect()
	if err != nil {
		log.Err(err)
	}

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

func createEntity(t *testing.T) entity.Entity {
	svc := entity.New(db)
	entity, err := svc.Create(entity.Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
		Items: []entity.EntityItems{entity.EntityItems{ItemName: "#1"}, entity.EntityItems{ItemName: "#2"}},
	})
	assert.Nil(t, err)
	assert.Greater(t, entity.ID, uint(0), "New ID ==0")
	return entity
}

func TestCreateEntity(t *testing.T) {
	ety := entity.Entity{
		Name:  "Entity X",
		Value: 43,
		Date:  time.Now(),
		Items: []entity.EntityItems{entity.EntityItems{ItemName: "#1"}, entity.EntityItems{ItemName: "#2"}},
	}
	url := fmt.Sprintf("%s/", baseURL)
	rr := executeReq(t, "POST", url, ety)
	assert.Equal(t, 200, rr.Code, "Return code not 200")
	var results []map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &results)
	assert.NotNil(t, results, "Entity is nil")
	assert.Greater(t, len(results), 0, "Error body does not equal 1")
}

func executeReq(t *testing.T, verb string, url string, body interface{}) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	var bodyStr string
	if body != struct{}{} {
		bdy, err := json.Marshal(body)
		if err != nil {
			t.Errorf("Error marshalling object: %s", err)
		} else {
			bodyStr = string(bdy)
		}
	}
	req, err := http.NewRequest(verb, url, strings.NewReader(bodyStr))
	if err != nil {
		t.Errorf("Error creating new request: %s", err)
	}
	router.ServeHTTP(rr, req)
	checkStatus(t, rr)
	return rr
}

func checkStatus(t *testing.T, rr *httptest.ResponseRecorder) {
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
