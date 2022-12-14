package auth

import "p2p/wallet"

type Success struct {
	UserID        string `json:"user_id"`
	Username      string `json:"username"`
	PhoneNumber   string `json:"phone_number"`
	wallet.Wallet `json:"wallet"`
	Fail          `json:"fail"`
}

type Fail struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
}
