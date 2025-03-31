package entities

type Master struct {
	Tag string
	Description string
	Unit string
	Component string
}

type ReadCSV struct {
	headers []string
	records []map[string]string
	err error
}