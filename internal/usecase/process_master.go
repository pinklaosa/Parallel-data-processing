package usecase

type MasterUsecase struct {
	repo CSVRepository
}

type CSVRepository interface {
	ReadCSV() ([]string, []map[string]string, error)
}

func NewMasterUsecase(repo CSVRepository) *MasterUsecase {
	return &MasterUsecase{repo: repo}
}

func (r *MasterUsecase) GroupingMaster() (map[string][]string, error) {
	headers, records, err := r.repo.ReadCSV()
	if err != nil {
		return nil, err
	}

	master := make(map[string][]string)
	for _, head := range headers {
		seen := make(map[string]struct{})
		var hearderName []string
		for _,rec := range records {
			if _,exists := seen[rec[head]]; !exists {	
				seen[rec[head]] = struct{}{}
				hearderName = append(hearderName, rec[head])
			}
		}
		master[head] = hearderName
	}
	return master, nil
}
