package transaction

type Transaction struct {
	ID        string `json:"id"`
	PostedAt  int64  `json:"postedAt"`
	AccountId string `json:"accountId"`
	Title     string `json:"title"`
	Amount    int64  `json:"amount"`
}

func (t Transaction) GetID() string {
	return t.ID
}
