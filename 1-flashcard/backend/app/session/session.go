package session

import "irayspace.com/flashcard/app/learnable"

type Session struct {
	ID            string                `json:"id"`
	Learnables    []learnable.Learnable `json:"learnables"`
	LastSessionAt int64                 `json:"lastSessionAt"`
}
