package learnable

import "irayspace.com/flashcard/app/card"

type Learnable struct {
	ID           string    `json:"id"`
	Card         card.Card `json:"card"`
	LastReviewAt int64     `json:"lastReviewAt"`
	Interval     int64     `json:"interval"`
}
