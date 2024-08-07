package tx

import (
	"cmp"
	"errors"
	"slices"
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
		ID:          uuid.NewString(),
		Amount:      dto.Amount,
		Currency:    "PHP",
		TxType:      dto.TxType,
		Description: dto.Description,
		Source:      dto.Source,
		Destination: dto.Destination,
		PostedAt:    time.Now().UTC().Unix(),
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
	} else if dto.TxType == "TRANSFER" {
		err := createTransferEntries(&tx)
		if err != nil {
			return &tx, err
		}
	}

	return &tx, nil
}

func GetTxsByNumber(number string) ([]domain.Tx, error) {
	txs, err := repository.GetAll[domain.Tx]("./tmp/transactions.txt")
	if err != nil {
		return nil, err
	}

	filteredTxs := []domain.Tx{}
	for _, tx := range txs {
		attachTxDescription(number, &tx)
		if tx.Source != nil && tx.Source.Number == number {
			filteredTxs = append(filteredTxs, tx)
		} else if tx.Destination != nil && tx.Destination.Number == number {
			filteredTxs = append(filteredTxs, tx)
		}
	}

	// Sort Z-A by PostedAt
	slices.SortFunc(filteredTxs, func(a, b domain.Tx) int {
		return cmp.Compare(b.PostedAt, a.PostedAt)
	})

	return filteredTxs, nil
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
		TxID:      tx.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

func createTransferEntries(tx *domain.Tx) error {
	err := ledgerentry.CreateLedgerEntry(&ledgerentry.LedgerEntry{
		Number:    tx.Source.Number,
		Amount:    tx.Amount,
		Currency:  tx.Currency,
		Direction: "CREDIT",
		TxID:      tx.ID,
	})
	if err != nil {
		return err
	}

	err = ledgerentry.CreateLedgerEntry(&ledgerentry.LedgerEntry{
		Number:    tx.Destination.Number,
		Amount:    tx.Amount,
		Currency:  tx.Currency,
		Direction: "DEBIT",
		TxID:      tx.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

// Should be in the creation of transactions
func attachTxDescription(number string, tx *domain.Tx) {
	if tx.Destination.Number == number && tx.Source != nil {
		tx.Description = "Received from " + tx.Source.Name
	} else if tx.Destination.Number == number && tx.Source == nil {
		tx.Description = "Deposited money"
	} else if tx.Source != nil && tx.Source.Number == number {
		tx.Description = "Sent to " + tx.Destination.Name
	}
}
