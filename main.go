package main

import "parallel-data-processing/infrastructure"

func main() {
	CSVInfrastructure := infrastructure.NewCSVReader("/assets/GTG-2_MASTER_DATA.csv")
	CSVInfrastructure.ReadCSV()
}