package tests

import (
	"database/sql"
	"walletTask/database"
	"walletTask/handlers"

	"github.com/gin-gonic/gin"
)

// тут подключаемся к БД для тестирования 
func StartTestDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://evans:evans@localhost:8001/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// создание тестовых маршрутов и тестируемых обработчиков
func StartTestServer(db *sql.DB) *gin.Engine {
	database.DB = db // нужно для того, чтобы обращение к database.DB в коде взаимодествовоало с БД к которой подключились выше 
	router := gin.Default()
	router.POST("api/v1/wallet/create", handlers.WalletCreate)
	router.POST("api/v1/wallet", handlers.WalletOperation)
	router.GET("api/v1/wallets/:WALLET_UUID", handlers.WalletGetBalance)
	return router
}
