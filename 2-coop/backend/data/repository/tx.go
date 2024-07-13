package repository

import (
	"encoding/json"
	"fmt"
	"os"

	"irayspace.com/coop/domain"
)

func SaveTx(a *domain.Tx) error {
	txs, err := GetAll[domain.Tx]("./tmp/transactions.txt")
	if err != nil {
		return err
	}

	file, err := os.Create("./tmp/transactions.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, tx := range txs {
		var data []byte
		if a.ID == tx.ID {
			data, err = json.Marshal(a)
		} else {
			data, err = json.Marshal(tx)
		}
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(file, "%s\n", data)
		if err != nil {
			return err
		}
	}

	return nil
}
