package transfer

type Reciver struct {
	Username            string `json:"username"` // this will be the email username
	DestinationWalletID string `json:"dst_wallet_id"`
}

type Sender struct {
	Username       string `json:"username"` // this will be the email username
	UserID         string `json:"user_id"`
	SourceWalletID string `json:"src_wallet_id"` // user_uid same for wallet_id
	Amount         int    `json:"amount"`
}

// transfer from bank to wallet
type TranferFounds struct {
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	WalletId string `json:"wallet_id"`
	Amount   int    `json:"amount"`
}

// New transfer intent
type Transfer struct {
	Sender        `json:"sender"`           // new transfer intent
	Reciver       `json:"reciver"`          // recive from sender
	TranferFounds `json:"transafer_founds"` // from bank to user wallet
}

// Transactions
type TransactionIntent struct {
	Sender  `json:"sender"`
	Reciver `json:"reciver"`
}
type Transaction struct {
	Proceed        bool   `json:"proceed"`
	TxAmountIntent int    `json:"tx_amount_intent"`
	TxStatusCode   int    `json:"tx_status_code"`
	TxMessage      string `json:"tx_message"`
}

type CardNotFound struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
}
