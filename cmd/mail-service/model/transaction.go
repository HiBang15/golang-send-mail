package model

import (
	db "github.com/HiBang15/golang-send-mail.git/db/mail/sqlc"
	"context"
	"database/sql"
	"fmt"
	"sync"
)

type Transaction struct {
	db *sql.DB
	*db.Queries
}

var transaction *Transaction
var once sync.Once

func NewTransaction(conn *sql.DB) *Transaction {
	once.Do(func() {
		trans = &Transaction{
			db:      conn,
			Queries: db.New(conn),
		}
	})
	return trans
}

func (trans *Transaction) execTx(ctx context.Context, fn func(queries *db.Queries) error) error{
	tx, err := trans.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)

	err = fn(q)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()

}
