package main

import "walletTask/server"


// @title           Test Task
// @version         0.1
// @description     Тестовое задание

// @contact.name   Evans Trein
// @contact.email  evanstrein@icloud.com
// @contact.url  https://github.com/EvansTrein

// @host      localhost:8000
// @schemes   http

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	server.StartServer()
}
