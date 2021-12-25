package wallet

import (
	"github.com/ahmetavc/wallet/pkg/infra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWalletCreate(t *testing.T) {
	wallet, err := NewWallet()

	assert.Nil(t, err)
	assert.Equal(t, wallet.balance, float64(0))
}

func TestWalletFromDTO(t *testing.T) {
	dto := infra.WalletDTO{
		Id:      "id",
		Balance: float64(10),
	}

	wallet := NewWalletFromDTO(dto)

	assert.Equal(t, wallet.id, dto.Id)
	assert.Equal(t, wallet.balance, dto.Balance)
}

func TestWalletDeposit(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Deposit(float64(100))

	assert.Nil(t, err)
	assert.Equal(t, wallet.balance, float64(110))
}

func TestWalletDepositWithZeroAmount(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Deposit(float64(0))

	assert.NotNil(t, err)
	assert.Equal(t, wallet.balance, float64(10))
}

func TestWalletDepositWithNegativeAmount(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Deposit(float64(-10))

	assert.NotNil(t, err)
	assert.Equal(t, wallet.balance, float64(10))
}

func TestWalletWithdraw(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Withdraw(float64(5))

	assert.Nil(t, err)
	assert.Equal(t, wallet.balance, float64(5))
}

func TestWalletWithdrawWhenNotEnough(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Withdraw(float64(15))

	assert.NotNil(t, err)
	assert.Equal(t, wallet.balance, float64(10))
}

func TestWalletWithdrawWithZeroAmount(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Withdraw(float64(0))

	assert.NotNil(t, err)
	assert.Equal(t, wallet.balance, float64(10))
}

func TestWalletWithdrawWithNegativeAmount(t *testing.T) {
	wallet := &wallet{
		id:      "uuid",
		balance: 10,
	}

	err := wallet.Withdraw(float64(-10))

	assert.NotNil(t, err)
	assert.Equal(t, wallet.balance, float64(10))
}