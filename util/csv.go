package util

import (
	"encoding/csv"
	"os"
)

// ReadCsvFile - Returns an array with the csv rows
func ReadCsvFile(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return make([][]string, 0), err
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// AppendToCsvFile - Appends row to filename csv
func AppendToCsvFile(filename string, row []string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	err = w.WriteAll([][]string{row})
	if err != nil {
		return err
	}

	return nil
}
