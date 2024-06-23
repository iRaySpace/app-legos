package card

type Card struct {
	ID    string `json:"id"`
	Front string `json:"front" validate:"required"`
	Back  string `json:"back" validate:"required"`
}
