package get_balance

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func (m *BalanceGatewayMock) GetBalance(account_id string) (*entity.Balance, error) {
	args := m.Called(account_id)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func TestGetBalanceUseCase_Execute(t *testing.T) {
	m := &BalanceGatewayMock{}
	m.On("GetBalance", mock.Anything).Return(32.0, nil)
	uc := NewGetBalanceUseCase(m)

	output, err := uc.Execute("account_id")
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, 32.0, output.Balance)
	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "GetBalance", 1)
}
