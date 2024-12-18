package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"walletTask/models"

	"github.com/stretchr/testify/assert"
)

func TestWalletGetBalance(t *testing.T) {
	testingUUID := "testUuid-c981-4a69-81c5-9a5a4838c78e"
	testingURL := fmt.Sprintf("/api/v1/wallets/%s", testingUUID)

	db, err := StartTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	router := StartTestServer(db)

	tetsReq, err := http.NewRequest("GET", testingURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("testWalletCreation", func(t *testing.T) {
		newWallet := models.NewWallet{Uuid: testingUUID}
		jsonData, err := json.Marshal(newWallet)
		if err != nil {
			t.Fatal(err)
		}

		tetsReqCreat, err := http.NewRequest("POST", "/api/v1/wallet/create", bytes.NewBuffer(jsonData))
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		router.ServeHTTP(w, tetsReqCreat)

		assert.Equal(t, http.StatusCreated, w.Code, w.Body.String())
	})

	t.Run("testGetWallet", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, tetsReq)

		assert.Equal(t, http.StatusOK, w.Code, w.Body.String())
	})

	t.Run("testGetWalletNoUUID", func(t *testing.T) {
		testingUUID := "123-123-123"
		testingURL := fmt.Sprintf("/api/v1/wallets/%s", testingUUID)

		tetsReq, err := http.NewRequest("GET", testingURL, nil)
		if err != nil {
			t.Fatal(err)
		}

		w := httptest.NewRecorder()
		router.ServeHTTP(w, tetsReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	deleteQuery := "DELETE FROM wallets WHERE walletID = $1"
	_, err = db.Exec(deleteQuery, testingUUID)
	if err != nil {
		log.Println("failed to delete the tested uuid ->>", err)
		return
	}

	log.Println("the wallet test is executed, the test UUID is deleted from the table")
}
