package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all function to execute ab queries and transactions
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// SQLStore provides all function to execute ab queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
} 

func NewStore(db *sql.DB) Store {
	return &SQLStore{db: db, Queries: New(db)}
}

// ExecTx executes a function within a database transaction
func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error{
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v ,rb error: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}
