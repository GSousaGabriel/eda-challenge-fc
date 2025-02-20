package database

import (
	"database/sql"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	if err != nil {
		panic(err)
	}

	return &AccountDB{
		DB: db,
	}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client
	account.Client = &client

	stmt, err := a.DB.Prepare("Select a.id, a.client_id, a.balance, a.created_at, c.id, c.name, c.email, c.created_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(id)
	err = row.Scan(
		&account.ID,
		&account.Client.ID,
		&account.Balance,
		&account.CreatedAt,
		&client.ID,
		&client.Name,
		&client.Email,
		&client.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	stmt, err := a.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (a *AccountDB) UpdateBalance(account *entity.Account) error {
	stmt, err := a.DB.Prepare("UPDATE accounts SET balance = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account.Balance, account.ID)
	if err != nil {
		return err
	}
	return nil
}
