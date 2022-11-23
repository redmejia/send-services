package dbpostgres

import (
	"context"
	newaccount "p2p/settingAccount"
	"strings"
	"time"
)

func (db *DBPostgres) InsertNewUser(user *newaccount.Register) bool {

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
		`insert into user_bank_info (user_uid, full_name, user_card, created_at)
			values ($1, $2, $3, $4)
		`, userUID, user.Bank.FullName, user.Bank.Card, time.Now())

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

	err = tx.Commit()
	if err != nil {
		db.ErrorLog.Fatal(err)
	}

	// all insertd
	return true
}
