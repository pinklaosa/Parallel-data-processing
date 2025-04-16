package usecase

import (
	"fmt"
	"maps"
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

// key is the time, value is the map of string to float64
type HourlyData map[time.Time]map[string]float64

func (r *RawUsecase) SamplingData(layoutDatetime string) (HourlyData, error) {
	_, records, err := r.repo.ReadCSV()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	numWorkers := 4
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
					timestamp, err := time.Parse(layoutDatetime, hourStr)
					hour := timestamp.Truncate(time.Hour)
					if err != nil {
						fmt.Println("Error parsing time:", err)
						break
					}
					if _,exists := hourly[hour]; !exists {
						hourly[hour] = make(map[string]float64)
						hourly[hour]["count"] = 0
					}
					for key, value := range rec {
						if key != "TimeStamp" {
							val, err := strconv.ParseFloat(value, 64)
							if err == nil {
								hourly[hour][key] += val
							}
						}
					}
					hourly[hour]["count"]++
				}
				result <- hourly
			}
		}()
	}

	go func() {
		for i := 0; i < len(records); i += chunkSize {
			end := min(i+chunkSize, len(records))
			chunkData <- records[i:end]
		}
		close(chunkData)
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	var allChunkData []HourlyData
	for res := range result {
		allChunkData = append(allChunkData, res)
	}

	// allchunkdata length is 4, so we need to merge them into one chunk
	mergedChunk := make(HourlyData)
	for _, chunk := range allChunkData {
		for hour, data := range chunk {
			if _, exists := mergedChunk[hour]; !exists {
				mergedChunk[hour] = make(map[string]float64)
			}
			maps.Copy(mergedChunk[hour], data)
		}
	}

	for _,data := range mergedChunk {
		count := data["count"]
		delete(data, "count")

		for k,sum := range data {
			data[k] = sum / count
		}
	}

	return mergedChunk, nil
}
