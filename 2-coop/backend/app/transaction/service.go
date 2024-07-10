package transaction

import (
	"time"

	"github.com/google/uuid"
	"irayspace.com/coop/app/account"
	"irayspace.com/coop/data"
)

func CreateTransaction(t *Transaction) error {
	t.ID = uuid.NewString()
	t.PostedAt = time.Now().Unix()

	err := data.Save("./tmp/transactions.txt", t, false)
	if err != nil {
		return err
	}

	a, err := data.FindById[account.Account]("./tmp/accounts.txt", t.AccountId)
	if err != nil {
		return err
	}

	a.Balance = a.Balance + t.Amount

	err = data.Save("./tmp/accounts.txt", a, true)
	if err != nil {
		return err
	}

	return nil
}
