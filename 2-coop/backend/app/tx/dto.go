package tx

import "irayspace.com/coop/domain"

type CreateTxDTO struct {
	Amount      int64             `json:"amount"`
	Currency    string            `json:"currency"`
	Description string            `json:"description"`
	TxType      string            `json:"txType"` // DEPOSIT, TRANSFER
	Source      *domain.TxAccount `json:"source"`
	Destination *domain.TxAccount `json:"destination"`
}
