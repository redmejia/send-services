package dbpostgres

import (
	"context"
	account "p2p/settingAccount"
	"p2p/wallet"
	"strings"
	"time"
)

func (db *DBPostgres) InsertNewUser(user *account.Register) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := db.Db.BeginTx(ctx, nil)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}
	defer tx.Rollback()

	var userUID string

	row := tx.QueryRowContext(ctx,
		`insert into register (full_name, email, phone, created_at)
			values ($1, $2, $3, $4)
		RETURNING user_uid
	`, user.FullName, user.Email, user.Phone, time.Now())

	err = row.Scan(&userUID)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	_, err = tx.ExecContext(ctx,
		`insert into signin (user_uid, email, password, created_at)
		values ($1, $2, $3, $4)
	`, userUID, user.Email, user.Password, time.Now())

	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	_, err = tx.ExecContext(ctx,
		`insert into user_bank_info (user_uid, full_name, user_card, cv_number, created_at)
			values ($1, $2, $3, $4, $5)
		`, userUID, user.Bank.FullName, user.Bank.Card, user.Bank.CvNumber, time.Now())

	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	username := strings.Split(user.Email, "@")[0]
	_, err = tx.ExecContext(ctx,
		`insert into wallet (wallet_id, username, balance, created_at)
			values ($1, $2, $3, $4)
		`, userUID, username, 0, time.Now())

	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	user.UserUID = userUID

	err = tx.Commit()
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	// all insertd
	return true
}

func (db *DBPostgres) GetInfoBank(userID string) account.Bank {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.Db.QueryRowContext(ctx,
		`select user_card, cv_number from user_bank_info where user_uid = $1`, userID)

	var bankInfo account.Bank

	err := row.Scan(&bankInfo.Card, &bankInfo.CvNumber)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	return bankInfo
}

func (db *DBPostgres) ckeckWalletBalance(walletId string) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var srcWalletBalance int

	row := db.Db.QueryRowContext(ctx,
		`select balance from wallet 
		where wallet_id = $1`, walletId)

	err := row.Scan(&srcWalletBalance)
	if err != nil {
		return 0, err
	}

	return srcWalletBalance, nil
}

func (db *DBPostgres) TransferFromBankToWallet(walletId string, amount int) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	srcWalletBalance, err := db.ckeckWalletBalance(walletId)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	if srcWalletBalance == 0 {

		_, err := db.Db.ExecContext(ctx,
			`update wallet
					set balance = $1
				where wallet_id = $2`, amount, walletId)

		if err != nil {
			db.ErrorLog.Fatal("0 unpdate", err)
		}

	} else {

		var newBalance int = srcWalletBalance + amount
		_, err := db.Db.ExecContext(ctx,
			`update wallet
					set balance = $1
				where wallet_id = $2`, newBalance, walletId)

		if err != nil {
			db.ErrorLog.Fatal("plus one update ", err)
		}

	}

}

func (db *DBPostgres) TransferWalletToWallet(userUID string, destinationWallet string, amount int) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// check balance on wallet is grather to the amount
	// if the current balance in wallet or sender has enougth then
	// source wallet
	srcWalletBalance, err := db.ckeckWalletBalance(userUID)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	// destination wallet
	dstBalanceWallet, err := db.ckeckWalletBalance(destinationWallet)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	if srcWalletBalance > amount {
		tx, err := db.Db.BeginTx(ctx, nil)
		if err != nil {
			db.ErrorLog.Fatal(err)
		}
		defer tx.Rollback()

		// decrement balance on the source wallet
		var newSrcBalance int = srcWalletBalance - amount
		_, err = tx.ExecContext(ctx,
			`update wallet 
				set balance = $1
			where wallet_id = $2`, newSrcBalance, userUID)
		if err != nil {
			db.ErrorLog.Fatal(err)
		}

		if dstBalanceWallet == 0 {

			_, err := tx.ExecContext(ctx,
				`update wallet
					set balance = $1
				where wallet_id = $2`, amount, destinationWallet)
			if err != nil {
				db.ErrorLog.Fatal(err)
			}

		} else {

			// increment the balance on the detination wallet
			var newDstBalance int = dstBalanceWallet + amount
			_, err = tx.ExecContext(ctx,
				`update wallet
					set balance = $1
				where wallet_id = $2`, newDstBalance, destinationWallet)
			if err != nil {
				db.ErrorLog.Fatal(err)
			}
		}

		err = tx.Commit()
		if err != nil {
			db.ErrorLog.Fatal(err)
		}

		return true
	}

	return false
}

// wallet trasanction records
func (db *DBPostgres) InsertTrxsRecordWalletToWallet(userID string, destinationWallet string, amount int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.Db.ExecContext(ctx,
		`insert into wallet_transactions (user_uid, from_wallet, to_wallet, amount) 
			values ($1, $2, $3, $4)`, userID, userID, destinationWallet, amount)

	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	return true

}

// user_id primary key on the tale is the user uuid also is the same for the wallet
func (db *DBPostgres) InsertTrxsRecordBankToWallet(userID string, bankCard string, amount int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// userID is same as wallet
	_, err := db.Db.ExecContext(ctx,
		`insert into bank_transactions (user_uid, from_bank, to_wallet, amount) 
			values ($1, $2, $3, $4)`, userID, bankCard, userID, amount)
	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	return true
}

func (db *DBPostgres) GetWalletInformation(walletID string) wallet.Wallet {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var myWallet wallet.Wallet

	row := db.Db.QueryRowContext(ctx,
		`select username, balance 
			from wallet 
		where wallet_id = $1`, walletID)

	err := row.Scan(&myWallet.Username, &myWallet.Balance)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	return myWallet

}
