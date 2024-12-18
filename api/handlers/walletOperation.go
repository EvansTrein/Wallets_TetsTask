package handlers

import (
	"log"
	"walletTask/database"
	"walletTask/models"

	"github.com/gin-gonic/gin"
)

// @Summary Perform a wallet operation
// @Description Perform a deposit or withdrawal operation on the wallet with the given walletID
// @Tags Wallets
// @Accept json
// @Produce json
// @Param wallet body models.WalletRequest true "Wallet operation request"
// @Success 200 {object} models.RespMessage
// @Failure 400 {object} models.RespError
// @Failure 500 {object} models.RespError
// @Router /api/v1/wallet [post]
func WalletOperation(ctx *gin.Context) {
	var walletReq models.WalletRequest // variable for incoming data

	// retrieving data from the request body
	if err := ctx.BindJSON(&walletReq); err != nil {
		log.Println("invalid data in the request body")
		ctx.JSON(400, models.RespError{MessageErr: "invalid data in the request body", TextErr: err.Error()})
		return
	}
	// transaction verification
	if walletReq.Operation != "DEPOSIT" && walletReq.Operation != "WITHDRAW" {
		log.Println("invalid operation in 'operationType'")
		ctx.JSON(400, models.RespError{MessageErr: "invalid operation in 'operationType'", TextErr: "None"})
		return
	}

	// Amount correctness check
	if walletReq.Amount <= 0 {
		log.Println("invalid number in 'Amount'")
		ctx.JSON(400, models.RespError{MessageErr: "invalid number in 'Amount'", TextErr: "number must be positive"})
		return
	}

	if walletReq.Operation == "DEPOSIT" {
		err := database.SqlDeposit(walletReq.WalletID, walletReq.Amount)
		if err != nil {
			switch err.Error() {
			case "there is no such wallet":
				ctx.JSON(400, models.RespError{MessageErr: "no wallet by passed walletId", TextErr: err.Error()})
				return
			default:
				ctx.JSON(500, models.RespError{MessageErr: "failed to execute SQL query (DEPOSIT) to the database", TextErr: err.Error()})
				return
			}
		}
	}

	if walletReq.Operation == "WITHDRAW" {
		err := database.SqlWithdraw(walletReq.WalletID, walletReq.Amount)
		if err != nil {
			switch err.Error() {
			case "there is no such wallet":
				ctx.JSON(400, models.RespError{MessageErr: "no wallet by passed walletId", TextErr: "None"})
				return
			case "insufficient funds in the wallet":
				ctx.JSON(400, models.RespError{MessageErr: "insufficient funds in the wallet", TextErr: "None"})
				return
			default:
				ctx.JSON(500, models.RespError{MessageErr: "failed to execute SQL query (WITHDRAW) to the database", TextErr: err.Error()})
				return
			}
		}
	}

	ctx.JSON(200, models.RespMessage{Message: "operation successful"})
}
