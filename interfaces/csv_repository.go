package interfaces

type CSVRepository struct {
	filePath string
	reader   CSVReader
}

type CSVReader interface {
	ReadCSV() ([][]string, error)
}

func NewCSVRepository(filepath string, reader CSVReader) *CSVRepository {
	return &CSVRepository{filePath: filepath, reader: reader}
}

func (c *CSVRepository) ReadCSV() ([][]string, error) {
	return c.reader.ReadCSV()
}