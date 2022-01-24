package main

import (
	"api-reference-golang/api"
	"api-reference-golang/events/kafka"
	"api-reference-golang/models"
	"api-reference-golang/models/migrate"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"os"
	"strconv"
)

// @title           Reference Golang API
// @version         1.0
// @description     Reference Golang API

// @contact.name   API Support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	username := osVariable("DB_USER", "postgres")
	password := osVariable("DB_PASS", "fixme")
	dbName := osVariable("DB_NAME", "entity")
	dbHost := osVariable("DB_HOST", "localhost")
	webPort := osVariable("PORT", "8080")
	kafkaUrl := osVariable("KAFKA_URL", "localhost:9092")
	debug, _ := strconv.ParseBool(osVariable("DEBUG", "false"))

	// Setup Logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// Set Debugging
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup Database
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	db := models.New(postgres.Open(dbUri))

	err := db.Connect()
	if err != nil {
		log.Panic().Err(err)
	}
	defer db.Disconnect()
	migrate.Migrate(db.DB())

	// Setup Kafka
	msg := kafka.New(kafkaUrl, "entity_events", "wawa/api-reference-golang", "")

	// Setup Router
	r := api.Router(db, msg)
	log.Printf("Starting server on the port http://0.0.0.0:XXXX...")
	err = r.Run(fmt.Sprintf(":%s", webPort))
	if err != nil {
		log.Panic().Err(err)
	}
}

func osVariable(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		val = def
	}
	return val
}
