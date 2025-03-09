package config

import (
	"factory-pattern/enums"
	"os"
	"sync"
)

type IConfig interface {
	GetFileStorageType() enums.FileStorageTypeEnum
}
type Config struct {
	fileStorageType enums.FileStorageTypeEnum
}

var env IConfig

const (
	fileStorageTypeKey = "fileStorageType"
)

func InitConfig() {
	val := os.Getenv(fileStorageTypeKey)
	fileStorageType := enums.Local_FileStorageType
	if val == string(enums.Local_FileStorageType) {
		fileStorageType = enums.Local_FileStorageType
	} else if val == string(enums.S3_FileStorageType) {
		fileStorageType = enums.S3_FileStorageType
	} else {
		panic("invalid file storage type")
	}
	env = &Config{
		fileStorageType: fileStorageType,
	}

}

func GetConfig() IConfig {
	if env == nil {
		sync.OnceFunc(func() { InitConfig() })()
	}
	return env
}

func (e *Config) GetFileStorageType() enums.FileStorageTypeEnum {
	return e.fileStorageType
}
