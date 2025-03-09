package service

type IFileService interface {
	GetFileUrl(fileName string) string
}
