package handlers

import (
	"strings"
	"walletTask/database"
	"walletTask/models"

	"github.com/gin-gonic/gin"
)

func WalletCreate(ctx *gin.Context) {
	var newWallet models.NewWallet

	if err := ctx.BindJSON(&newWallet); err != nil {
		ctx.JSON(400, models.RespError{MessageErr: "invalid data in the request body", TextErr: err.Error()})
		return
	}

	if err := database.SqlCreateWallet(newWallet.Uuid); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			ctx.JSON(400, models.RespError{MessageErr: "wallet with this UUID already exists", TextErr: err.Error()})
			return
		}
		ctx.JSON(500, models.RespError{MessageErr: "failed to create a wallet in the database", TextErr: err.Error()})
		return
	}

	ctx.JSON(201, models.RespMessage{Message: "the wallet has been successfully created"})
}
