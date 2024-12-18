package database

import (
	"database/sql"
	"errors"
	"log"
	"walletTask/models"
)

func SqlCreateWallet(walletId string) error {
	sqlComm := "INSERT INTO wallets (walletid, total) VALUES ($1, $2)"

	_, err := DB.Exec(sqlComm, walletId, 0)
	if err != nil {
		return err
	}

	return nil
}

func SqlDeposit(walletId string, amount float64) error {
	var exists bool // variable for checking the existence of a record in the database

	// initiate a transaction
	transaction, err := DB.Begin()
	if err != nil {
		log.Println("failed to create a transaction")
		return err
	}

	// pass the SQL command and block the record until the transaction is committed or rolled back
	err = transaction.QueryRow("SELECT EXISTS(SELECT 1 FROM wallets WHERE walletID = $1 FOR UPDATE)", walletId).Scan(&exists)
	if err != nil {
		log.Println("failed to search the database by walletID")
		transaction.Rollback()
		return err
	}

	// check availability
	if !exists {
		transaction.Rollback()
		return errors.New("there is no such wallet")
	}

	// pass the SQL command to update
	_, err = transaction.Exec("UPDATE wallets SET total = total + $1 WHERE walletID = $2", amount, walletId)
	if err != nil {
		log.Println("failed to update the amount")
		transaction.Rollback()
		return err
	}

	// time.Sleep(time.Second * 10)

	// finalize the transaction, saving the changes
	err = transaction.Commit()
	if err != nil {
		log.Println("failed to save transaction changes")
		transaction.Rollback()
		return err
	}

	return nil
}

func SqlWithdraw(walletId string, amount float64) error {
	var activeWallet models.ActiveWallet // variable for the wallet we are working with

	// initiate a transaction
	transaction, err := DB.Begin()
	if err != nil {
		log.Println("failed to create a tranzation")
		return err
	}

	// pass the SQL command and block the record until the transaction is committed or rolled back
	err = transaction.QueryRow("SELECT walletid, total FROM wallets WHERE walletID = $1 FOR UPDATE",
		walletId).Scan(&activeWallet.Uuid, &activeWallet.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			transaction.Rollback()
			return errors.New("there is no such wallet")
		}
		log.Println("failed to search the database by walletID")
		transaction.Rollback()
		return err
	}

	// time.Sleep(time.Second * 10)

	// verification that there are sufficient funds for the operation
	if amount > activeWallet.Total {
		transaction.Rollback()
		return errors.New("insufficient funds in the wallet")
	}

	// pass the SQL command to update
	_, err = transaction.Exec("UPDATE wallets SET total = total - $1 WHERE walletID = $2 AND total >= $1", amount, walletId)
	if err != nil {
		log.Println("failed to update the amount")
		transaction.Rollback()
		return err
	}

	// finalize the transaction, saving the changes
	err = transaction.Commit()
	if err != nil {
		log.Println("failed to save transaction changes")
		transaction.Rollback()
		return err
	}

	return nil
}

func SqlGetWallet(walletId string, wallet *models.ActiveWallet) error {

	errResultFind := DB.QueryRow("SELECT walletid, total FROM wallets WHERE walletID = $1", walletId).Scan(&wallet.Uuid, &wallet.Total)
	if errResultFind != nil {
		if errResultFind == sql.ErrNoRows {
			return errors.New("there is no such wallet")
		}
		return errResultFind
	}

	return nil
}
