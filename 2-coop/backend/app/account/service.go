package account

import (
	"github.com/google/uuid"
	"irayspace.com/coop/data"
)

func CreateAccount(a *Account) error {
	a.ID = uuid.NewString()

	err := data.Save("./tmp/accounts.txt", a, false)
	if err != nil {
		return err
	}

	return nil
}
