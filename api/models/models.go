package models

type WalletRequest struct {
	WalletID  string  `json:"walletId" binding:"required"`
	Operation string  `json:"operationType" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

type RespError struct {
	MessageErr string
	TextErr    string
}

type RespMessage struct {
	Message string
}

type NewWallet struct {
	Uuid string `json:"walletId" binding:"required"`
}

type ActiveWallet struct {
	Uuid   string
	Total  float64
}

