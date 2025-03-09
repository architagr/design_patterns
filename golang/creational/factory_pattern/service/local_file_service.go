package service

import "fmt"

type localFileService struct {
	basePath   string
	resellerId int
}

func InitLocalFileService(resellerId int, baseFolderPath string) IFileService {
	return &localFileService{
		basePath:   baseFolderPath,
		resellerId: resellerId,
	}
}
func (svc *localFileService) GetFileUrl(fileName string) string {

	return fmt.Sprintf("%s://%s/%s", svc.getSftpUrl(), svc.basePath, fileName)
}

func (svc *localFileService) getSftpUrl() string {
	return fmt.Sprintf("local-sftp-%d", svc.resellerId)
}
