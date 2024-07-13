package domain

type Account struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Number    string `json:"number"`
	Balance   int64  `json:"balance"`
	Currency  string `json:"currency"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func (a Account) GetID() string {
	return a.ID
}
