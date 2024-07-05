package data

import (
	"encoding/json"
	"fmt"
	"os"
)

func Save(f string, v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(f, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprintf(file, "%s\n", data); err != nil {
		return err
	}

	return nil
}
