package usecase

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type RawUsecase struct {
	repo   CSVRepository
	master map[string][]string
}

type RawRepository interface {
	ReadCSV() ([]string, []map[string]string, error)
}

func NewRawUsecase(repo CSVRepository, master map[string][]string) *RawUsecase {
	return &RawUsecase{repo: repo, master: master}
}

func (r *RawUsecase) GroupingRaw() {
	headers, _, err := r.repo.ReadCSV()
	if err != nil {
		return
	}
	fmt.Println("Headers:", headers)
}

type HourlyData map[time.Time]map[string]float64

func (r *RawUsecase) SamplingData() {
	_, records, err := r.repo.ReadCSV()
	if err != nil {
		return
	}
	var wg sync.WaitGroup
	numWorkers := 5
	chunkSize := (len(records) + numWorkers - 1) / numWorkers
	chunkData := make(chan []map[string]string, numWorkers)
	result := make(chan HourlyData, numWorkers)

	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for chunk := range chunkData {
				hourly := make(HourlyData)
				for _, rec := range chunk {
					hourStr := rec["TimeStamp"]
					
					hour, err := time.Parse(time.RFC3339, hourStr)
					if err != nil {
						continue
					}
					datum := make(map[string]float64)
					for key, value := range rec {
						if key != "TimeStamp" {
							val, err := strconv.ParseFloat(value, 64)
							if err == nil {
								datum[key] = val
							}
						}
					}
					hourly[hour] = datum
				}
				result <- hourly
			}
		}()
	}

	go func() {
		for range numWorkers {
			for i := 0; i < len(records); i += chunkSize {
				end := i + chunkSize
				if end > len(records) {
					end = len(records)
				}
				chunkData <- records[i:end]
			}
		}
		close(chunkData)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	var mergedData []HourlyData
	for res := range result {
		mergedData = append(mergedData, res)
	}
	fmt.Println("Merged Data:",mergedData)

}
