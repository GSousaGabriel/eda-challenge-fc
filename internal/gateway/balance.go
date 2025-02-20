package gateway

import "github.com.br/devfullcycle/fc-ms-wallet/internal/entity"

type BalanceGateway interface {
	GetBalance(account_id string) (*entity.Balance, error)
}
