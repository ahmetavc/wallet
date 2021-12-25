package wallet

import (
	"errors"
	"github.com/google/uuid"
)

type Wallet struct {
	Id      string  `json:"Id"`
	Balance float64 `json:"Balance"`
}

func NewWallet() (*Wallet, error){
	uuid, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	return &Wallet{
		Id:      uuid.String(),
		Balance: 0,
	}, nil
}

func (w *Wallet) GetBalance() float64{
	return w.Balance
}

func (w *Wallet) GetId() string{
	return w.Id
}

func (w *Wallet) Withdraw(amount float64) error {
	if amount <= 0{
		return errors.New("withdrawal amount cannot be zero or negative")
	}

	if amount > w.Balance {
		return errors.New("not enough Balance")
	}

	w.Balance -= amount

	return nil
}

func (w *Wallet) Deposit(amount float64) error {
	if amount <= 0{
		return errors.New("deposit amount cannot be zero or negative")
	}

	w.Balance += amount

	return nil
}



