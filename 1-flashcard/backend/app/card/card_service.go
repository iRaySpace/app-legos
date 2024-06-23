package card

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func CreateCard(card *Card) error {
	card.ID = uuid.NewString()

	data, err := json.Marshal(card)
	if err != nil {
		return nil
	}

	file, err := os.OpenFile("./tmp/cards.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return nil
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "%s\n", data); err != nil {
		return nil
	}

	return nil
}
