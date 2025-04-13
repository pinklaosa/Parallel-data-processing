package main

import (
	"parallel-data-processing/internal/infrastructure"
	"parallel-data-processing/internal/interfaces"
	"parallel-data-processing/internal/usecase"
)

// [Infrastructure] → [Repository] → [Use Case] → [Presenter] → [Output]
func main() {
	csvReaderMaster := infrastructure.NewCSVReader("./assets/GTG-1_MASTER_DATA.csv")
	csvInterfaceMaster := interfaces.NewCSVRepository(csvReaderMaster);
	masterUsecase := usecase.NewMasterUsecase(csvInterfaceMaster)
	groupMaster,_ := masterUsecase.GroupingMaster()
	//raw data is the large data csv
	csvReaderRaw := infrastructure.NewCSVReader("./assets/GTG-1.csv")
	csvInterfaceRaw := interfaces.NewCSVRepository(csvReaderRaw);
	rawData := usecase.NewRawUsecase(csvInterfaceRaw,groupMaster)
	rawData.GroupingRaw()
	//next to use case 

}