package account

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"irayspace.com/coop/data/repository"
	"irayspace.com/coop/domain"
)

func CreateAccount(dto *CreateAccountDTO) (*domain.Account, error) {
	a := domain.Account{
		ID:        uuid.NewString(),
		CreatedAt: time.Now().UTC().Unix(),
		UpdatedAt: time.Now().UTC().Unix(),
		Name:      dto.Name,
		Number:    dto.Number,
		Balance:   0,
		Currency:  "PHP",
	}

	err := repository.Save("./tmp/accounts.txt", a)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func GetAccounts() ([]domain.Account, error) {
	return repository.GetAll[domain.Account]("./tmp/accounts.txt")
}

func GetAccount(number string) (*domain.Account, error) {
	accounts, err := repository.GetAll[domain.Account]("./tmp/accounts.txt")
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		if account.Number == number {
			return &account, nil
		}
	}
	return nil, errors.New("account not found")
}
