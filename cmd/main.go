package main

import (
	"parallel-data-processing/internal/infrastructure"
)

// [Infrastructure] → [Repository] → [Use Case] → [Presenter] → [Output]
func main() {
	filePath := infrastructure.NewCSVReader("./assets/GTG-2_MASTER_DATA.csv")
	filePath.ReadCSV()
}
