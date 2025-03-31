package main

import (
	"fmt"
	"strategy_pattern/enums"
	exportstrategy "strategy_pattern/export_strategy"
	"strategy_pattern/models"
	"strategy_pattern/persistence"
	"strategy_pattern/service"
)

func main() {
	userPersistenceObj := persistence.UserPersistence{}

	userSvc := service.InitUserService(&userPersistenceObj)

	// Json Builder
	jsonBuilder, err := exportstrategy.UserExportBuilder(&models.ExportConfig{
		Type: enums.ExportType_Json,
	})
	if err != nil {
		panic(err)
	}
	data, err := userSvc.ExportUsers(models.Pagination{}, jsonBuilder)
	fmt.Println("error : \t", err)
	fmt.Println("data json: \t", data)

	// CSV Builder
	csvBuilder, err := exportstrategy.UserExportBuilder(&models.ExportConfig{
		Type: enums.ExportType_CSV,
	})
	if err != nil {
		panic(err)
	}
	data, err = userSvc.ExportUsers(models.Pagination{}, csvBuilder)
	fmt.Println("error : \t", err)
	fmt.Println("data csv: \t", data)

	// XML Builder
	xmlBuilder, err := exportstrategy.UserExportBuilder(&models.ExportConfig{
		Type: enums.ExportType_Xml,
	})
	if err != nil {
		panic(err)
	}
	data, err = userSvc.ExportUsers(models.Pagination{}, xmlBuilder)
	fmt.Println("error : \t", err)
	fmt.Println("data xml: \t", data)
}
