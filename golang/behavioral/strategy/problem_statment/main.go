package main

import (
	"fmt"
	"strategy/enums"
	"strategy/models"
	"strategy/persistence"
	"strategy/service"
)

func main() {
	userPersistenceObj := persistence.UserPersistence{}

	userSvc := service.InitUserService(&userPersistenceObj)

	data, err := userSvc.ExportUsers(models.Pagination{}, models.ExportConfig{
		Type: enums.ExportType_Json,
	})
	fmt.Println("error : \t", err)
	fmt.Println("data json: \t", data)

	data, err = userSvc.ExportUsers(models.Pagination{}, models.ExportConfig{
		Type: enums.ExportType_CSV,
	})
	fmt.Println("error : \t", err)
	fmt.Println("data csv: \t", data)

	data, err = userSvc.ExportUsers(models.Pagination{}, models.ExportConfig{
		Type: enums.ExportType_Xml,
	})
	fmt.Println("error : \t", err)
	fmt.Println("data xml: \t", data)

}
