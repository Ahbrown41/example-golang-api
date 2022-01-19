package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"log"
	"os"
	"reference-golang-svc/models"
	"reference-golang-svc/router"
)

// @title Reference Golang API
// @version 1.0
// @description This API is an example on using Golang for an APIO

// @host localhost:27017
// @BasePath /api/v1
func main() {
	username := osVariable("db_user", "postgres")
	password := osVariable("db_pass", "fixme")
	dbName := osVariable("db_name", "public")
	dbHost := osVariable("db_host", "localhost")

	// Setup Database
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	conn := models.New(postgres.Open(dbUri))

	err := conn.Connect()
	if err != nil {
		log.Panic(err)
	}
	defer conn.Disconnect()

	models.Migrate(conn.DB())

	// Setup Router
	r := router.Router()
	log.Printf("Starting server on the port http://0.0.0.0:XXXX...")
	r.Run()
}

func osVariable(name string, def string) string {
	val := os.Getenv(name)
	if val == "" {
		val = def
	}
	return val
}
