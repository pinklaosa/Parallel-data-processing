package main

import (
	"parallel-data-processing/internal/infrastructure"
	"parallel-data-processing/internal/interfaces"
	"parallel-data-processing/internal/usecase"
)

// [Infrastructure] → [Repository] → [Use Case] → [Presenter] → [Output]
func main() {
	csvReader := infrastructure.NewCSVReader("./assets/GTG-1_MASTER_DATA.csv")
	csvInterface := interfaces.NewCSVRepository(csvReader);
	masterUsecase := usecase.NewMasterUsecase(csvInterface)
	masterUsecase.GroupingMaster()
	//next to use case 

}