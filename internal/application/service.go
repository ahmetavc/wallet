package application

import (
	"github.com/ahmetavc/wallet/internal/domain/wallet"
	"github.com/ahmetavc/wallet/pkg/infra"
)

type Repository interface {
	Upsert(key string, value *infra.WalletDTO) error
	Get(key string) (infra.WalletDTO, error)
	Remove(key string) error
	Close() error
}

type service struct {
	Repo Repository
}

func NewService(repository Repository) *service {
	return &service{Repo: repository}
}

func (s *service) Create() (string, error) {
	w, _ := wallet.NewWallet()

	dto := &infra.WalletDTO{
		Id:      w.GetId(),
		Balance: w.GetBalance(),
	}

	err := s.Repo.Upsert(dto.Id, dto)

	if err != nil {
		return "", err
	}

	return w.GetId(), nil
}

func (s *service) Get(id string) (float64, error) {
	walletDTO, err := s.Repo.Get(id)

	w := wallet.NewWalletFromDTO(walletDTO)

	if err != nil {
		return float64(0), err
	}

	return w.GetBalance(), nil
}

func (s *service) Deposit(id string, amount float64) (float64, error) {
	walletDTO, err := s.Repo.Get(id)

	w := wallet.NewWalletFromDTO(walletDTO)

	err = w.Deposit(amount)

	if err != nil {
		return 0, err
	}

	newWalletDTO := &infra.WalletDTO{
		Id:      w.GetId(),
		Balance: w.GetBalance(),
	}

	err = s.Repo.Upsert(id, newWalletDTO)

	if err != nil {
		return 0, err
	}

	return newWalletDTO.Balance, nil
}

func (s *service) Withdraw(id string, amount float64) (float64, error) {
	walletDTO, err := s.Repo.Get(id)

	w := wallet.NewWalletFromDTO(walletDTO)

	err = w.Withdraw(amount)

	if err != nil {
		return 0, err
	}

	newWalletDTO := &infra.WalletDTO{
		Id:      w.GetId(),
		Balance: w.GetBalance(),
	}

	err = s.Repo.Upsert(id, newWalletDTO)

	if err != nil {
		return 0, err
	}

	return newWalletDTO.Balance, nil
}