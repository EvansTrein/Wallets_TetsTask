package server

import (
	"walletTask/envs"
	"walletTask/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "walletTask/docs"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()

	router.POST("api/v1/wallet/create", handlers.WalletCreate)

	router.POST("api/v1/wallet", handlers.WalletOperation)

	router.GET("api/v1/wallets/:WALLET_UUID", handlers.WalletGetBalance)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":" + envs.ServerEnvs.API_PORT)
}
