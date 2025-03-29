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

func (c *CSVReader) ReadCSV() {
	file, err := os.Open(c.filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var wg sync.WaitGroup
	reader := csv.NewReader(bufio.NewReader(file))
	// header,_ := reader.Read()
	rows := make(chan []string)
	// var records sync.Map
	
	workers := 4

	go func() {
		for {
			row, err := reader.Read()
			if err == io.EOF {
				close(rows)
			}
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			rows <- row
		}
	}()

	//convert struct
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
		}()
	}
	

}
