package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// TODO: handle existing object
func Save(f string, v any, e bool) error {
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

type HasID interface {
	GetID() string
}

func FindById[T HasID](f string, id string) (*T, error) {
	file, err := os.OpenFile(f, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		var row T
		if err := json.Unmarshal(sc.Bytes(), &row); err != nil {
			return nil, err
		}
		if row.GetID() == id {
			return &row, nil
		}
	}

	return nil, nil
}
