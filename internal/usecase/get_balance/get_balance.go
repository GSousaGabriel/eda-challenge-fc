package get_balance

import (
	"github.com.br/devfullcycle/fc-ms-wallet/internal/gateway"
)

type GetBalanceOutputDTO struct {
	Balance float64
}

type GetBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(balanceGateway gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		BalanceGateway: balanceGateway,
	}
}

func (uc *GetBalanceUseCase) Execute(account_id string) (*GetBalanceOutputDTO, error) {
	balance, err := uc.BalanceGateway.GetBalance(account_id)

	if err != nil {
		return nil, err
	}

	output := &GetBalanceOutputDTO{
		Balance: balance.Balance,
	}
	return output, nil
}
