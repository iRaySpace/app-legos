package ledgerentry

type LedgerEntry struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Direction string `json:"direction"` // DEBIT, CREDIT
	TxID      string `json:"txId"`
	PostedAt  int64  `json:"postedAt"`
}
