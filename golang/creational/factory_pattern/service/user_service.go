package service

import (
	"factory-pattern/config"
	"factory-pattern/enums"
	"factory-pattern/models"
	"fmt"
)

type IUserService interface {
	GetUser(resellerId, id int) (*models.GetUserResponse, error)
}

type userService struct {
	envVariables config.IConfig
}

var userServiceObj IUserService

func InitUserService(envVariables config.IConfig) IUserService {
	if userServiceObj == nil {
		userServiceObj = &userService{envVariables: envVariables}
	}
	return userServiceObj
}
func (svc *userService) GetUser(resellerId, id int) (*models.GetUserResponse, error) {
	var fileSvc IFileService
	if svc.envVariables.GetFileStorageType() == enums.Local_FileStorageType {
		fileSvc = InitLocalFileService(resellerId, "user")
	} else if svc.envVariables.GetFileStorageType() == enums.S3_FileStorageType {
		awsConfig := InitAwsService()
		fileSvc = InitS3FileService(awsConfig, resellerId, "user")
	}
	return &models.GetUserResponse{
		UserId:    id,
		UserName:  "test User",
		AvatarUrl: fileSvc.GetFileUrl(fmt.Sprintf("avatar%d.jpg", id)),
	}, nil
}
