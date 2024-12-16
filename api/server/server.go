package server

import (
	"log"
	"walletTask/database"
	"walletTask/envs"
)

func StartServer() {
	if err := envs.LoadEnvs(); err != nil {
		log.Fatalln("Не удалось загрузить env файл\n ошибка ->", err.Error())
	}

	if err := database.InitDatabase(); err != nil {
		log.Fatalln("Не удалось подключиться к базе данных\n ошибка ->", err.Error())
	}

	InitRoutes()
}