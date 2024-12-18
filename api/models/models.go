package models

type WalletRequest struct {
	WalletID  string  `json:"walletId" binding:"required" example:"38c7e784-e963-4cc1-9124-de3e6c7e60e4"`
	Operation string  `json:"operationType" binding:"required" example:"DEPOSIT or WITHDRAW"`
	Amount    float64 `json:"amount" binding:"required" example:"500"`
}

type RespError struct {
	MessageErr string `example:"user error"`
	TextErr    string `example:"error text"`
}

type RespMessage struct {
	Message string `example:"message"`
}

type NewWallet struct {
	Uuid string `json:"walletId" binding:"required" example:"38c7e784-e963-4cc1-9124-de3e6c7e60e4"`
}

type ActiveWallet struct {
	Uuid  string  `example:"38c7e784-e963-4cc1-9124-de3e6c7e60e4"`
	Total float64 `example:"1000"`
}
