package handlers

import (
	"walletTask/database"
	"walletTask/models"

	"github.com/gin-gonic/gin"
)

func WalletGetBalance(ctx *gin.Context) {
	uuidWallet := ctx.Param("WALLET_UUID")
	var activeWallet models.ActiveWallet

	if uuidWallet == "" {
		ctx.JSON(400, models.RespError{MessageErr: "WALLET_UUUID has not been transmitted", TextErr: "None"})
		return
	}

	if err := database.SqlGetWallet(uuidWallet, &activeWallet); err != nil {
		switch err.Error() {
		case "there is no such wallet":
			ctx.JSON(400, models.RespError{MessageErr: "no wallet by passed walletId", TextErr: err.Error()})
			return
		default:
			ctx.JSON(500, models.RespError{MessageErr: "failed to execute SQL query to the database", TextErr: err.Error()})
			return
		}
	}

	ctx.JSON(200, activeWallet)
}
