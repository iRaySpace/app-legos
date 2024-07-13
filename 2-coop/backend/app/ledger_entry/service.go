package ledgerentry

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"irayspace.com/coop/data/repository"
)

func CreateLedgerEntry(le *LedgerEntry) error {
	le.ID = uuid.NewString()
	le.PostedAt = time.Now().Unix()

	err := repository.Save("./tmp/ledger_entries.txt", le)
	if err != nil {
		return err
	}

	err = setBalance(le)
	if err != nil {
		return err
	}

	return nil
}

func setBalance(le *LedgerEntry) error {
	a, err := repository.FindAccountByNumber(le.Number)
	if err != nil {
		return err
	}
	if a == nil {
		return errors.New("account not found")
	}

	if le.Direction == "DEBIT" {
		a.Balance = a.Balance + le.Amount
	} else if le.Direction == "CREDIT" {
		a.Balance = a.Balance - le.Amount
	}

	err = repository.SaveAccount(a)
	if err != nil {
		return err
	}

	return nil
}
