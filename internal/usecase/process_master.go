package usecase

import (
	"fmt"
	"sync"
)

type MasterUsecase struct {
	repo CSVRepository
}

type CSVRepository interface {
	ReadCSV() ([]string, []map[string]string, error)
}

func NewMasterUsecase(repo CSVRepository) *MasterUsecase {
	return &MasterUsecase{repo: repo}
}

func (r *MasterUsecase) GroupingMaster() {
	headers, records, err := r.repo.ReadCSV()
	if err != nil {
		fmt.Println(err)
	}

	var wg sync.WaitGroup
	var datum map[string][]string
	for _, g := range headers {
	}

	workers := 3

	recordsChan := make(chan map[string]string, len(records))

	for range workers {
		go func() {
			var g []string
			for rec := range recordsChan {
				for _, h := range headers {
					m.Store(h, append(g, rec[h]))
				}
			}
		}()
	}

	go func() {
		for _, rec := range records {
			recordsChan <- rec
		}
	}()

	go func() {
		wg.Wait()
		close(recordsChan)
	}()

	m.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})

}
