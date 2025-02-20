package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type TransactionDB struct {
	DB *sql.DB
}

func NewTransactionDB(db *sql.DB) *TransactionDB {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")
	if err != nil {
		panic(err)
	}

	return &TransactionDB{
		DB: db,
	}
}

func (t *TransactionDB) Create(transaction *entity.Transaction) error {
	stmt, err := t.DB.Prepare("INSERT INTO transactions (id, account_id_from, account_id_to, amount, created_at) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.ID, transaction.AccountFrom.ID, transaction.AccountTo.ID, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
