package database

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/config"
	_ "github.com/lib/pq"
)

func GetConnectionString(cfg config.Config) string {
	dbUser := cfg.Database.Username
	dbPass := cfg.Database.Password
	dbHost := cfg.Database.Host
	dbPort := cfg.Database.Port
	dbName := cfg.Database.Name

	connection := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	return connection
}

func HealthCheckDb(cfg config.Config) {
	connectionString := GetConnectionString(cfg)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to db succeeded!")
}