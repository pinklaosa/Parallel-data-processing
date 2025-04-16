package main

import (
	"fmt"
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

	// [raw] -> [mean hr] -> [linear regression] -> [mean sd]
	csvReaderRaw := infrastructure.NewCSVReader("./assets/GTG-1.csv")
	csvInterfaceRaw := interfaces.NewCSVRepository(csvReaderRaw);
	rawData := usecase.NewRawUsecase(csvInterfaceRaw,groupMaster);

	_,err := rawData.SamplingData("2006-01-02 15:04:05-07:00")
	if err != nil {
		fmt.Println(err)
	}
	

}