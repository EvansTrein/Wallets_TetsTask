package server

import (
	"walletTask/envs"
	"walletTask/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.POST("api/v1/wallet/create", handlers.WalletCreate)

	router.POST("api/v1/wallet", handlers.WalletOperation)

	router.GET("api/v1/wallets/:WALLET_UUID", handlers.WalletGetBalance)

	router.Run(":" + envs.ServerEnvs.API_PORT)
}
