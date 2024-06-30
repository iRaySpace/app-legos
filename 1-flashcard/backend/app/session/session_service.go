package session

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func CreateSession(s *Session) error {
	s.ID = uuid.NewString()

	s.LastSessionAt = time.Now().Unix()

	for i := range s.Learnables {
		s.Learnables[i].LastReviewAt = time.Now().Unix()
		s.Learnables[i].Interval = s.Learnables[i].Interval + 86400 // TODO: for change on values
	}

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
