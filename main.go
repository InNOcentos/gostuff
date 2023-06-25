package main

import (
	"myapp/config"
	"myapp/server"
	"myapp/database"
)



func main() {
	cfg := config.Config{}
	config.LoadEnv(&cfg)

	database.HealthCheckDb(cfg)

	server := http.InitServer()

	server.Bootstrap(cfg)
}
