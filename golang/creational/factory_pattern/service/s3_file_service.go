package service

import "fmt"

type s3FileService struct {
	bucketName string
	resellerId int
	apiConfig  *AwsService
}

func InitS3FileService(apiConfig *AwsService, resellerId int, bucketName string) IFileService {
	return &s3FileService{
		bucketName: bucketName,
		resellerId: resellerId,
		apiConfig:  apiConfig,
	}
}

func (svc *s3FileService) GetFileUrl(fileName string) string {
	return fmt.Sprintf("s3://%s/%d/%s", svc.bucketName, svc.resellerId, fileName)
}
