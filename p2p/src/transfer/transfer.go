package transfer

// New trian
type TransferIntent struct {
	UserID            string `json:"user_id"` // user_uid same for wallet_id
	DestinationWallet string `json:"dst_wallet"`
	Amount            int    `json:"amount"`
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
