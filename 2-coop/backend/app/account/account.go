package account

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhoneNo   string `json:"phoneNo"`
	Balance   int64  `json:"balance"`
}

func (a Account) GetID() string {
	return a.ID
}
