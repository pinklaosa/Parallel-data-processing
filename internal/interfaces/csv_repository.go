package interfaces

type CSVRepository struct {
	reader CSVReader
}

type CSVReader interface {
	ReadCSV() ([]string, []map[string]string, error)
}

func NewCSVRepository(reader CSVReader) *CSVRepository {
	return &CSVRepository{reader: reader}
}

func (c *CSVRepository) ReadCSV() ([]string, []map[string]string, error) {
	return c.reader.ReadCSV()
}
