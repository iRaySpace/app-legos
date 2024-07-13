package domain

type TxAccount struct {
	Number string `json:"number"`
	Name   string `json:"name"`
}

type Tx struct {
	ID          string     `json:"id"`
	Amount      int64      `json:"amount"`
	Currency    string     `json:"currency"`
	Description string     `json:"description"`
	TxType      string     `json:"txType"`
	Source      *TxAccount `json:"source,omitempty"`
	Destination *TxAccount `json:"destination,omitempty"`
	PostedAt    int64      `json:"postedAt"`
}

func (t Tx) GetID() string {
	return t.ID
}
