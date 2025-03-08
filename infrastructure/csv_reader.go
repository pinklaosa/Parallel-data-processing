package infrastructure

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	filePath string
}

func NewCSVReader(filePath string) *CSVReader {
	return &CSVReader{filePath: filePath}
}

func (c *CSVReader) ReadCSV() ([][]string, error) {
	file, err := os.Open(c.filePath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		file.Close()
		return nil, err
	}
	return record, nil
}
