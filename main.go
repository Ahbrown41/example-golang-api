package main

import (
	"api-reference-golang/api"
	"api-reference-golang/models"
	"api-reference-golang/models/migrate"
	"fmt"
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
	username := osVariable("db_user", "postgres")
	password := osVariable("db_pass", "fixme")
	dbName := osVariable("db_name", "entity")
	dbHost := osVariable("db_host", "localhost")
	webPort := osVariable("PORT", "8080")
	debug, _ := strconv.ParseBool(osVariable("DEBUG", "false"))

	// Setup Logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Info().Msg("Service starting")

	// Setup Database
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn := models.New(postgres.Open(dbUri))

	err := conn.Connect()
	if err != nil {
		log.Panic().Err(err)
	}
	defer conn.Disconnect()

	migrate.Migrate(conn.DB())

	// Setup Router
	r := api.Router(conn)
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
