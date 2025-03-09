package controller

import (
	"factory-pattern/logger"
	"factory-pattern/models"
	"factory-pattern/service"
	"fmt"
)

type IUserController interface {
	GetUser(resellerId, id int) (*models.GetUserResponse, error)
}

type userController struct {
	service service.IUserService
	logger  logger.ILogger
}

var userControllerObj IUserController

func GetUserController() (IUserController, error) {
	if userControllerObj == nil {
		return nil, fmt.Errorf("user Controller not initilized")
	}
	return userControllerObj, nil
}
func InitJobController(userServiceObj service.IUserService, logObj logger.ILogger) IUserController {
	if userControllerObj == nil {
		userControllerObj = &userController{
			service: userServiceObj,
			logger:  logObj,
		}
	}
	return userControllerObj
}

func (ctlr *userController) GetUser(resellerId, id int) (*models.GetUserResponse, error) {
	return ctlr.service.GetUser(resellerId, id)
}
