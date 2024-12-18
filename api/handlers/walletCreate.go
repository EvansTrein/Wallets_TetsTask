package handlers

import (
	"strings"
	"walletTask/database"
	"walletTask/models"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new wallet
// @Description Create a new wallet with zero balance and specified UUID
// @Tags Wallets
// @Accept json
// @Produce json
// @Param wallet body models.NewWallet true "Wallet object"
// @Success 201 {object} models.RespMessage
// @Failure 400 {object} models.RespError
// @Failure 500 {object} models.RespError
// @Router /api/v1/wallet/create [post]
func WalletCreate(ctx *gin.Context) {
	var newWallet models.NewWallet // variable for incoming data

	// retrieving data from the request body
	if err := ctx.BindJSON(&newWallet); err != nil {
		ctx.JSON(400, models.RespError{MessageErr: "invalid data in the request body", TextErr: err.Error()})
		return
	}

	// SQL command call
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
