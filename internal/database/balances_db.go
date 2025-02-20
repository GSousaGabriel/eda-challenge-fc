package database

import (
	"database/sql"
	"fmt"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type BalancesDB struct {
	DB *sql.DB
}

func NewBalancesDB(db *sql.DB) *BalancesDB {
	return &BalancesDB{
		DB: db,
	}
}

func (a *BalancesDB) GetBalance(AccountId string) (*entity.Balance, error) {
	var balance entity.Balance

	stmt, err := a.DB.Prepare("SELECT balance FROM accounts WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(AccountId)
	fmt.Println("dçaljdasçk143298dsak")
	fmt.Println(AccountId)
	fmt.Println(row)
	err = row.Scan(&balance.Balance)

	if err != nil {
		return nil, err
	}
	return &balance, nil
}
