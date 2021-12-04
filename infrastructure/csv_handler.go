package infrastructure

import (
	"encoding/csv"
	"os"
)

type CsvHandler struct {
	filename string
}

func NewCsvHandler(filename string) *CsvHandler {
	return &CsvHandler{filename: filename}
}

// ReadCsvFile - Returns an array with the csv rows
func (csvHandler *CsvHandler) Read() ([][]string, error) {
	file, err := os.Open(csvHandler.filename)
	if err != nil {
		return make([][]string, 0), err
	}

	reader := csv.NewReader(file)
	return reader.ReadAll()
}

// AppendToCsvFile - Appends row to filename csv
func (csvHandler *CsvHandler) Append(row []string) error {
	file, err := os.OpenFile(csvHandler.filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
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
