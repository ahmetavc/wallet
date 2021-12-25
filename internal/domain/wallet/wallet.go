package wallet

import (
	"errors"
	"github.com/google/uuid"
)

type wallet struct {
	id string
	balance float64
}

func NewWallet() (*wallet, error){
	uuid, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &wallet{
		id: uuid.String(),
		balance: 0,
	}, nil
}

func (w *wallet) GetBalance() float64{
	return w.balance
}

func (w *wallet) Withdraw(amount float64) error {
	if amount <= 0{
		return errors.New("withdrawal amount cannot be zero or negative")
	}

	if amount > w.balance{
		return errors.New("not enough balance")
	}

	w.balance -= amount

	return nil
}

func (w *wallet) Deposit(amount float64) error {
	if amount <= 0{
		return errors.New("deposit amount cannot be zero or negative")
	}

	w.balance += amount

	return nil
}



