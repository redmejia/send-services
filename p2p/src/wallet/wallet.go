package wallet

type Wallet struct {
	UserID    string `json:"user_id"`
	WalletId  string `json:"wallet_id"` // uu_id
	Username  string `json:"username"`  // this will be the email username
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}
