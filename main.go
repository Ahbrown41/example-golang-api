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

// @title Reference Golang API
// @version 1.0
// @description This API is an example on using Golang for an API

// @host localhost:27017
// @BasePath /api/v1
func main() {
	username := osVariable("db_user", "postgres")
	password := osVariable("db_pass", "fixme")
	dbName := osVariable("db_name", "public")
	dbHost := osVariable("db_host", "localhost")
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
		log.Err(err)
		log.Panic()
	}
	defer conn.Disconnect()

	migrate.Migrate(conn.DB())

	// Setup Router
	r := api.Router()
	log.Printf("Starting server on the port http://0.0.0.0:XXXX...")
	err = r.Run()
	if err != nil {
		log.Err(err)
		log.Panic()
	}
}

func osVariable(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		val = def
	}
	return val
}
