package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strategy/enums"
	"strategy/models"
	"strings"
)

var (
	Err_NotSupportedExportType = errors.New("Not Supported export Type")
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

func (userSvc *UserService) ExportUsers(paginationObj models.Pagination, exportConfig models.ExportConfig) (string, error) {
	// Get users from the Store based on the pagination/filter selected by the user
	users := userSvc.GetUsers(paginationObj)

	// validate the type in which user wants the data in
	// and then convert the data in the required format
	switch exportConfig.Type {
	case enums.ExportType_Json:
		jsonData, err := json.Marshal(users)
		return string(jsonData), err
	case enums.ExportType_CSV:
		sb := strings.Builder{}
		sb.WriteString("id,name,email\n")
		for _, u := range users {
			sb.WriteString(fmt.Sprintf("%d,%s,%s\n", u.Id, u.Name, u.Email))
		}
		return sb.String(), nil
	default:
		return "", Err_NotSupportedExportType
	}
}
