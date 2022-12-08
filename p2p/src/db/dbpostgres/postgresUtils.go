package dbpostgres

import (
	"context"
	"p2p/auth"
	"p2p/transfer"
	"p2p/wallet"
	"strings"
	"time"
)

func (db *DBPostgres) InsertNewUser(user *auth.Register) bool {

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

	username := strings.Split(user.Email, "@")[0]

	_, err = tx.ExecContext(ctx,
		`insert into user_bank_info (user_uid, username, full_name, user_card, cv_number, created_at)
			values ($1, $2, $3, $4, $5, $6)
		`, userUID, username, user.Bank.FullName, user.Bank.Card, user.Bank.CvNumber, time.Now())

	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	_, err = tx.ExecContext(ctx,
		`insert into wallet (user_uid, username, balance, created_at)
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

func (db *DBPostgres) GetInfoBank(userID string) auth.Bank {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.Db.QueryRowContext(ctx,
		`select user_card, cv_number from user_bank_info where user_uid = $1`, userID)

	var bankInfo auth.Bank

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

// Need the wallet id not user id I will pass as paramert later
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

func (db *DBPostgres) TransferWalletToWallet(sender *transfer.Sender, reciver *transfer.Reciver) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// check balance on wallet is grather to the amount
	// if the current balance in wallet or sender has enougth then
	// source wallet
	srcWalletBalance, err := db.ckeckWalletBalance(sender.SourceWalletID)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	// destination wallet
	dstBalanceWallet, err := db.ckeckWalletBalance(reciver.DestinationWalletID)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	if srcWalletBalance > sender.Amount {
		tx, err := db.Db.BeginTx(ctx, nil)
		if err != nil {
			db.ErrorLog.Fatal(err)
		}
		defer tx.Rollback()

		// decrement balance on the source wallet
		var newSrcBalance int = srcWalletBalance - sender.Amount
		_, err = tx.ExecContext(ctx,
			`update wallet 
				set balance = $1
			where user_uid = $2 and wallet_id = $3`, newSrcBalance, sender.UserID, sender.SourceWalletID)
		if err != nil {
			db.ErrorLog.Fatal(err)
		}

		if dstBalanceWallet == 0 {
			// reciver does not share user_id but does share usename
			_, err := tx.ExecContext(ctx,
				`update wallet
					set balance = $1
				where username = $2 and wallet_id = $3`, sender.Amount, reciver.Username, reciver.DestinationWalletID)
			if err != nil {
				db.ErrorLog.Fatal(err)
			}

		} else {

			// increment the balance on the detination wallet
			var newDstBalance int = dstBalanceWallet + sender.Amount
			_, err = tx.ExecContext(ctx,
				`update wallet
					set balance = $1
				where username = $2 and wallet_id = $3`, newDstBalance, reciver.Username, reciver.DestinationWalletID)
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
func (db *DBPostgres) InsertTrxsRecordWalletToWallet(sender *transfer.Sender, reciver *transfer.Reciver) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.Db.ExecContext(ctx,
		`insert into wallet_transactions (user_uid, from_wallet, to_wallet, amount) 
			values ($1, $2, $3, $4)`, sender.SourceWalletID, sender.SourceWalletID, reciver.DestinationWalletID, sender.Amount)

	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	return true

}

// user_id primary key on the tale is the user uuid also is the same for the wallet
func (db *DBPostgres) InsertTrxsRecordBankToWallet(userID string, bankCard string, walletId string, amount int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// userID is same as wallet
	_, err := db.Db.ExecContext(ctx,
		`insert into bank_transactions (user_uid, from_bank, to_wallet, amount) 
			values ($1, $2, $3, $4)`, userID, bankCard, walletId, amount)
	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	return true
}

// my wallet information by wallet id
func (db *DBPostgres) GetWalletInformation(userID string) wallet.Wallet {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var myWallet wallet.Wallet

	row := db.Db.QueryRowContext(ctx,
		`select user_uid, wallet_id, username, balance, created_at
			from wallet 
		where user_uid = $1`, userID)

	err := row.Scan(&myWallet.UserID, &myWallet.WalletId, &myWallet.Username, &myWallet.Balance, &myWallet.CreatedAt)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	return myWallet

}

// For share wallet info username and wallet id to sender
func (db *DBPostgres) ShareWalletInfo(shareid string) wallet.Wallet {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var share wallet.Wallet

	row := db.Db.QueryRowContext(ctx,
		`select username, wallet_id 
		from wallet 
	where share_id = $1`, shareid)

	err := row.Scan(&share.Username, &share.WalletId)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	return share
}

func (db *DBPostgres) GetAuthSuccess(email string) auth.Success {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.Db.QueryRowContext(ctx,
		`select s.user_uid, w.username, r.phone from signin as s 
			join wallet as w on s.user_uid = w.user_uid
			join register as r on w.user_uid = r.user_uid
		where s.email = $1`, email)

	var success auth.Success
	err := row.Scan(&success.UserID, &success.Username, &success.PhoneNumber)
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	return success

}

func (db *DBPostgres) GetUserAuthInfo(email string, dstHashPassword *string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := db.Db.QueryRowContext(ctx,
		`select password from signin
			where email = $1 `, email)

	var pwd string
	err := row.Scan(&pwd)
	if err != nil {
		db.ErrorLog.Fatal(err)
		return false
	}

	*dstHashPassword = pwd

	return true
}
