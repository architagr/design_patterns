package service

import (
	"factory-pattern/config"
	"factory-pattern/enums"
	"factory-pattern/models"
	"fmt"
)

type IProductService interface {
	GetProduct(resellerId, id int) (*models.GetProductResponse, error)
}

type productService struct {
	envVariables config.IConfig
}

var productServiceObj IProductService

func InitProductService(envVariables config.IConfig) IProductService {
	if productServiceObj == nil {
		productServiceObj = &productService{envVariables: envVariables}
	}
	return productServiceObj
}
func (svc *productService) GetProduct(resellerId, id int) (*models.GetProductResponse, error) {
	var fileSvc IFileService
	if svc.envVariables.GetFileStorageType() == enums.Local_FileStorageType {
		fileSvc = InitLocalFileService(resellerId, "product")
	} else if svc.envVariables.GetFileStorageType() == enums.S3_FileStorageType {
		awsConfig := InitAwsService()
		fileSvc = InitS3FileService(awsConfig, resellerId, "product")
	}
	return &models.GetProductResponse{
		ProductId:    id,
		ProductName:  "test User",
		ProductImage: fileSvc.GetFileUrl(fmt.Sprintf("product%d.jpg", id)),
	}, nil
}
