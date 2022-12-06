package wallet

type Wallet struct {
	ShareID   string `json:"share_id"` // share this id to the sender
	UserID    string `json:"user_id"`
	WalletId  string `json:"wallet_id"` // wallet id
	Username  string `json:"username"`  // this will be the email username
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}
