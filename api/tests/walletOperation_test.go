package tests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"walletTask/models"

	"github.com/stretchr/testify/assert"
)

func TestWalletOperation(t *testing.T) {
	testingURL := "/api/v1/wallet"
	testingUUID := "testUuid-c981-4a69-81c5-9a5a4838c78e"
	var walletRequest models.WalletRequest

	db, err := StartTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	router := StartTestServer(db)

	t.Run("walletCreation", func(t *testing.T) {
		newWallet := models.NewWallet{Uuid: testingUUID}
		jsonData, err := json.Marshal(newWallet)
		if err != nil {
			t.Fatal(err)
		}

		testReqCreat, err := http.NewRequest("POST", "/api/v1/wallet/create", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReqCreat)

		assert.Equal(t, http.StatusCreated, w.Code, w.Body.String())
	})

	t.Run("walletOperationDeposit", func(t *testing.T) {
		walletRequest.WalletID = testingUUID
		walletRequest.Operation = "DEPOSIT"
		walletRequest.Amount = 2000.00
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		tetsReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, tetsReq)

		assert.Equal(t, http.StatusOK, w.Code, w.Body.String())
	})

	t.Run("walletOperationWithdraw", func(t *testing.T) {
		walletRequest.WalletID = testingUUID
		walletRequest.Operation = "WITHDRAW"
		walletRequest.Amount = 2000.00
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		testReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReq)

		assert.Equal(t, http.StatusOK, w.Code, w.Body.String())
	})

	t.Run("walletOperationInsufficientFunds", func(t *testing.T) {
		walletRequest.WalletID = testingUUID
		walletRequest.Operation = "WITHDRAW"
		walletRequest.Amount = 3000.00
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		tetsReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, tetsReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	t.Run("walletInvalidOperationType", func(t *testing.T) {
		walletRequest.WalletID = testingUUID
		walletRequest.Operation = "LOL"
		walletRequest.Amount = 10.00
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		tetsReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, tetsReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	t.Run("walletOperationNoUUID", func(t *testing.T) {
		walletRequest.WalletID = "123-123-123"
		walletRequest.Operation = "DEPOSIT"
		walletRequest.Amount = 1000.00
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		testReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	t.Run("walletOperationAmountZero", func(t *testing.T) {
		walletRequest.WalletID = testingUUID
		walletRequest.Operation = "DEPOSIT"
		walletRequest.Amount = 0
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		testReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	t.Run("walletOperationAmountNegative", func(t *testing.T) {
		walletRequest.WalletID = testingUUID
		walletRequest.Operation = "DEPOSIT"
		walletRequest.Amount = -1000.00
		jsonData, err := json.Marshal(walletRequest)
		if err != nil {
			t.Fatal(err)
		}

		testReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	deleteQuery := "DELETE FROM wallets WHERE walletID = $1"
	_, err = db.Exec(deleteQuery, testingUUID)
	if err != nil {
		log.Println("failed to delete the tested uuid ->>", err)
		return
	}

	log.Println("tests on operations are executed, test UUID is deleted from the table")
}
