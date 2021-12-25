package application

import (
	"github.com/ahmetavc/wallet/internal/domain/wallet"
	"github.com/ahmetavc/wallet/pkg/infra"
)

type service struct {
	Repo infra.Repository
}

func NewService(repository infra.Repository) *service {
	return &service{Repo: repository}
}

func (s *service) Create() (string, error) {
	w, _ := wallet.NewWallet()

	err := s.Repo.Upsert(w.GetId(), w)

	if err != nil {
		return "", err
	}

	return w.GetId(), nil
}

func (s *service) Get(id string) (float64, error) {
	w, err := s.Repo.Get(id)

	if err != nil {
		return float64(0), err
	}

	return w.GetBalance(), nil
}

func (s *service) Deposit(id string, amount float64) error {
	w, err := s.Repo.Get(id)

	err = w.Deposit(amount)

	if err != nil {
		return err
	}

	err = s.Repo.Upsert(id, w)

	if err != nil {
		return err
	}

	return nil
}

func (s *service) Withdraw(id string, amount float64) error {
	w, err := s.Repo.Get(id)

	err = w.Withdraw(amount)

	if err != nil {
		return err
	}

	err = s.Repo.Upsert(id, w)

	if err != nil {
		return err
	}

	return nil
}