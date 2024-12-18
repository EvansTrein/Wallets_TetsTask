package handlers

import (
	"walletTask/database"
	"walletTask/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get wallet balance
// @Description Get the balance of the wallet with the given UUID
// @Tags Wallets
// @Accept json
// @Produce json
// @Param WALLET_UUID path string true "UUID of the wallet" example(38c7e784-e963-4cc1-9124-de3e6c7e60e4)
// @Success 200 {object} models.ActiveWallet
// @Failure 400 {object} models.RespError
// @Failure 500 {object} models.RespError
// @Router /api/v1/wallets/{WALLET_UUID} [get]
func WalletGetBalance(ctx *gin.Context) {
	uuidWallet := ctx.Param("WALLET_UUID") // get the UUID from the request
	var activeWallet models.ActiveWallet   // variable for the data to be returned in the response

	// check that the UUID has been transferred
	if uuidWallet == "" {
		ctx.JSON(400, models.RespError{MessageErr: "WALLET_UUUID has not been transmitted", TextErr: "None"})
		return
	}

	// SQL command call
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
