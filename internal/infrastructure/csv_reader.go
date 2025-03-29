package infrastructure

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
)

type CSVReader struct {
	filePath string
}

func NewCSVReader(filePath string) *CSVReader {
	return &CSVReader{filePath: filePath}
}

func (c *CSVReader) ReadCSV() []map[string]string {
	file, err := os.Open(c.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	reader := csv.NewReader(bufio.NewReader(file))
	header, _ := reader.Read()
	rows := make(chan []string)
	records := make(chan map[string]string)
	data := make(chan []map[string]string)

	workers := 5

	//convert struct
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for row := range rows {
				rec := make(map[string]string)
				for i, col := range header {
					rec[col] = row[i]
				}
				records <- rec
			}
		}()
	}
	go func() {
		wg.Wait()
		close(records)
	}()

	go func() {
		for {
			row, err := reader.Read()
			if err == io.EOF {
				close(rows)
				break
			}
			if err != nil {
				continue
			}
			rows <- row
		}
	}()

	go func() {
		var destructureData []map[string]string
		for record := range records {
			destructureData = append(destructureData, record)
		}
		data <- destructureData
	}()

	wg.Wait()

	final := <-data
	fmt.Println(final)
	return final
}
