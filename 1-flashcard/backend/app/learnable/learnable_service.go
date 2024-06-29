package learnable

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	"irayspace.com/flashcard/app/card"
)

func CreateLearnable(card *card.Card) (*Learnable, error) {
	file, err := os.OpenFile("./tmp/learnables.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	learnable := Learnable{
		ID:           uuid.NewString(),
		Card:         *card,
		LastReviewAt: 0,
		Interval:     0,
	}

	data, err := json.Marshal(learnable)
	if err != nil {
		return nil, err
	}

	if _, err := fmt.Fprintf(file, "%s\n", data); err != nil {
		return nil, err
	}

	return &learnable, err
}

func GetLearnables() ([]Learnable, error) {
	file, err := os.OpenFile("./tmp/learnables.txt", os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	learnables := []Learnable{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var learnable Learnable
		if err := json.Unmarshal(sc.Bytes(), &learnable); err != nil {
			return nil, err
		}
		learnables = append(learnables, learnable)
	}

	return learnables, nil
}
