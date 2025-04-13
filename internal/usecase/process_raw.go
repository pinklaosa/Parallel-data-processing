package usecase

import "fmt"

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

func (r *RawUsecase) SamplingData(){
	
}