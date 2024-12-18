package server

import (
	"log"
	"walletTask/database"
	"walletTask/envs"
)

func StartServer() {
	if err := envs.LoadEnvs(); err != nil {
		log.Fatalln("Failed to load env file\n error ->", err.Error())
	}

	if err := database.InitDatabase(); err != nil {
		log.Fatalln("Failed to connect to the database\n error ->", err.Error())
	}

	InitRoutes()
}