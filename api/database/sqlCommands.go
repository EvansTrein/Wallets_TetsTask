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
	var exists bool

	transaction, err := DB.Begin()
	if err != nil {
		log.Println("не удалось создать транзакцию")
		return err
	}

	err = transaction.QueryRow("SELECT EXISTS(SELECT 1 FROM wallets WHERE walletID = $1)", walletId).Scan(&exists)
	if err != nil {
		log.Println("не удалось выполнить поиск в БД по walletID")
		transaction.Rollback()
		return err
	}

	if !exists {
		transaction.Rollback()
		return errors.New("there is no such wallet")
	}

	_, err = transaction.Exec("UPDATE wallets SET total = total + $1 WHERE walletID = $2", amount, walletId)
	if err != nil {
		log.Println("не удалось обновить сумму")
		transaction.Rollback()
		return err
	}

	err = transaction.Commit()
	if err != nil {
		log.Println("не удалось сохранить изменения транзакции")
		transaction.Rollback()
		return err
	}

	return nil
}

func SqlWithdraw(walletId string, amount float64) error {
	var activeWallet models.ActiveWallet

	transaction, err := DB.Begin()
	if err != nil {
		log.Println("Не удалось создать транзацию")
		return err
	}

	err = transaction.QueryRow("SELECT walletid, total FROM wallets WHERE walletID = $1 FOR UPDATE", walletId).Scan(&activeWallet.Uuid, &activeWallet.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			transaction.Rollback()
			return errors.New("there is no such wallet")
		}
		log.Println("не удалось выполнить поиск в БД по walletID")
		transaction.Rollback()
		return err
	}

	if amount > activeWallet.Total {
		transaction.Rollback()
		return errors.New("insufficient funds in the wallet")
	}

	_, err = transaction.Exec("UPDATE wallets SET total = total - $1 WHERE walletID = $2 AND total >= $1", amount, walletId)
	if err != nil {
		log.Println("не удалось обновить сумму")
		transaction.Rollback()
		return err
	}

	// time.Sleep(time.Second * 10)

	err = transaction.Commit()
	if err != nil {
		log.Println("не удалось сохранить изменения транзакции")
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
