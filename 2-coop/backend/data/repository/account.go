package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"irayspace.com/coop/domain"
)

func FindAccountByNumber(number string) (*domain.Account, error) {
	file, err := os.OpenFile("./tmp/accounts.txt", os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var row domain.Account
		if err := json.Unmarshal(sc.Bytes(), &row); err != nil {
			return nil, err
		}
		if row.Number == number {
			return &row, nil
		}
	}

	return nil, nil
}

func SaveAccount(a *domain.Account) error {
	accounts, err := GetAll[domain.Account]("./tmp/accounts.txt")
	if err != nil {
		return err
	}

	file, err := os.Create("./tmp/accounts.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, account := range accounts {
		var data []byte
		if a.ID == account.ID {
			data, err = json.Marshal(a)
		} else {
			data, err = json.Marshal(account)
		}
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(file, "%s\n", data)
		if err != nil {
			return err
		}
	}

	return nil
}
