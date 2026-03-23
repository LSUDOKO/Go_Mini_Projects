package main

import (
	"encoding/csv"
	"os"
)

func LoadRecipient(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	records,err:=r.ReadAll()
	
}
