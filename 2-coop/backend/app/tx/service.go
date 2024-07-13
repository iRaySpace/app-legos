package tx

import (
	"errors"
	"time"

	"github.com/google/uuid"
	ledgerentry "irayspace.com/coop/app/ledger_entry"
	"irayspace.com/coop/data/repository"
	"irayspace.com/coop/domain"
)

func CreateTx(dto *CreateTxDTO) (*domain.Tx, error) {
	err := validateTxFields(dto)
	if err != nil {
		return nil, err
	}

	tx := domain.Tx{
		ID:              uuid.NewString(),
		Amount:          dto.Amount,
		Currency:        "PHP",
		TxType:          dto.TxType,
		Description:     dto.Description,
		Source:          dto.Source,
		Destination:     dto.Destination,
		Balance:         0,
		BalanceCurrency: "PHP",
		PostedAt:        time.Now().Unix(),
	}
	err = repository.Save("./tmp/transactions.txt", tx)
	if err != nil {
		return nil, err
	}

	if dto.TxType == "DEPOSIT" {
		err := createDepositEntry(&tx)
		if err != nil {
			return &tx, err
		}
		err = setDepositBalance(&tx)
		if err != nil {
			return &tx, err
		}
	}

	return &tx, nil
}

func validateTxFields(dto *CreateTxDTO) error {
	if dto.Amount == 0 {
		return errors.New("amount should be set")
	}
	if dto.TxType == "" {
		return errors.New("txType should be set")
	}
	if dto.TxType == "DEPOSIT" && dto.Destination == nil {
		return errors.New("destination should be set")
	}
	if dto.TxType == "TRANSFER" && (dto.Destination == nil || dto.Source == nil) {
		return errors.New("source and/or destination should be set")
	}
	return nil
}

func createDepositEntry(tx *domain.Tx) error {
	err := ledgerentry.CreateLedgerEntry(&ledgerentry.LedgerEntry{
		Number:    tx.Destination.Number,
		Amount:    tx.Amount,
		Currency:  tx.Currency,
		Direction: "DEBIT",
	})
	if err != nil {
		return err
	}
	return nil
}

func setDepositBalance(tx *domain.Tx) error {
	a, err := repository.FindAccountByNumber(tx.Destination.Number)
	if err != nil {
		return err
	}
	tx.Balance = a.Balance
	repository.SaveTx(tx)
	return nil
}