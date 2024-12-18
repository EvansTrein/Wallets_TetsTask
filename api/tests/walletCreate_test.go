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

func TestWalletCreate(t *testing.T) {
	testingURL := "/api/v1/wallet/create"
	testingUUID := "testUuid-c981-4a69-81c5-9a5a4838c78e"

	db, err := StartTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	router := StartTestServer(db)

	newWallet := models.NewWallet{Uuid: testingUUID}
	jsonData, err := json.Marshal(newWallet)
	if err != nil {
		t.Fatal(err)
	}

	testReq, err := http.NewRequest("POST", testingURL, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	t.Run("walletCreation", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReq)

		assert.Equal(t, http.StatusCreated, w.Code, w.Body.String())
	})

	t.Run("duplicateWalletCreation", func(t *testing.T) {
		w := httptest.NewRecorder()
		router.ServeHTTP(w, testReq)

		assert.Equal(t, http.StatusBadRequest, w.Code, w.Body.String())
	})

	t.Run("invalidData", func(t *testing.T) {
		invalidWalletData := models.NewWallet{Uuid: ""}
		jsonData, err = json.Marshal(invalidWalletData)
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

	log.Println("tests for wallet creation are completed, test UUID is deleted from the table")
}
