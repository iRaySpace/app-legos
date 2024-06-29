package session

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func CreateSession(s *Session) error {
	s.ID = uuid.NewString()

	// TODO: update learnable values

	file, err := os.OpenFile("./tmp/sessions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintf(file, "%s\n", data); err != nil {
		return err
	}

	return nil
}
