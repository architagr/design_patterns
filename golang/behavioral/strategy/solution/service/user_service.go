package service

import (
	exportstrategy "strategy_pattern/export_strategy"
	"strategy_pattern/models"
)

type IUserPersistence interface {
	ListUsers(paginationObj models.Pagination) []models.UserModel
}
type UserService struct {
	userPersistence IUserPersistence
}

func InitUserService(userPersistence IUserPersistence) *UserService {
	return &UserService{
		userPersistence: userPersistence,
	}
}

func (userSvc *UserService) GetUsers(paginationObj models.Pagination) []models.UserModel {
	return userSvc.userPersistence.ListUsers(paginationObj)
}

func (userSvc *UserService) ExportUsers(paginationObj models.Pagination, exportConfig exportstrategy.IUserExport) (string, error) {
	// Get users from the Store based on the pagination/filter selected by the user
	users := userSvc.GetUsers(paginationObj)

	// validate the type in which the user wants the data in
	// and then convert the data into the required format
	return exportConfig.Export(users)
}
